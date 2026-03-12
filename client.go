package bosskg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	DefaultTimeout = 15 * time.Second
	TestGatewayURL = "http://testgateway.serviceshare.com/testapi/clientapi/clientBusiness/common"
)

type Client struct {
	baseURL string
	merID   string

	desKey    []byte
	merPriKey *PrivateKey
	youfuPub  *PublicKey

	restyClient *resty.Client
	requestID   func() string
	userAgent   string
	debug       bool
}

type Config struct {
	BaseURL            string
	MerchantID         string
	DESKey             string
	MerchantPrivateKey string
	YoufuPublicKey     string
}

type Option func(*Client)

func WithRestyClient(rc *resty.Client) Option {
	return func(c *Client) {
		c.restyClient = rc
	}
}

func WithRequestIDFunc(f func() string) Option {
	return func(c *Client) {
		c.requestID = f
	}
}

func WithUserAgent(ua string) Option {
	return func(c *Client) {
		c.userAgent = ua
	}
}

func WithDebug(debug bool) Option {
	return func(c *Client) {
		c.debug = debug
	}
}

func NewClient(cfg Config, opts ...Option) (*Client, error) {
	if cfg.BaseURL == "" {
		return nil, errors.New("baseURL is required")
	}
	if cfg.MerchantID == "" {
		return nil, errors.New("merchantID is required")
	}
	if cfg.DESKey == "" {
		return nil, errors.New("desKey is required")
	}
	if cfg.MerchantPrivateKey == "" {
		return nil, errors.New("merchantPrivateKey is required")
	}
	if cfg.YoufuPublicKey == "" {
		return nil, errors.New("youfuPublicKey is required")
	}

	merPri, err := ParsePrivateKey(cfg.MerchantPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("parse merchant private key: %w", err)
	}
	youfuPub, err := ParsePublicKey(cfg.YoufuPublicKey)
	if err != nil {
		return nil, fmt.Errorf("parse youfu public key: %w", err)
	}
	desKey, err := DeriveDESKey(cfg.DESKey)
	if err != nil {
		return nil, fmt.Errorf("derive des key: %w", err)
	}

	c := &Client{
		baseURL:     cfg.BaseURL,
		merID:       cfg.MerchantID,
		desKey:      desKey,
		merPriKey:   merPri,
		youfuPub:    youfuPub,
		restyClient: resty.New().SetTimeout(DefaultTimeout),
		requestID:   DefaultRequestID,
		userAgent:   "bosskg-go/0.1",
	}
	for _, opt := range opts {
		if opt != nil {
			opt(c)
		}
	}
	if c.restyClient == nil {
		c.restyClient = resty.New().SetTimeout(DefaultTimeout)
	}
	if c.requestID == nil {
		c.requestID = DefaultRequestID
	}
	return c, nil
}

func (c *Client) Do(ctx context.Context, funCode FunCode, req any, out any) (*RespMessage, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	reqBodyPlain, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal reqData: %w", err)
	}

	if c.debug {
		log.Printf("[bosskg-debug] >>> ReqData (plain): %s", string(reqBodyPlain))
	}

	reqDataCipherB64, err := EncryptDESBase64(reqBodyPlain, c.desKey)
	if err != nil {
		return nil, fmt.Errorf("encrypt reqData: %w", err)
	}

	sign, err := SignRSASHA1Base64([]byte(reqDataCipherB64), c.merPriKey)
	if err != nil {
		return nil, fmt.Errorf("sign reqData: %w", err)
	}

	msg := &ReqMessage{
		ReqID:   c.requestID(),
		FunCode: funCode,
		MerID:   c.merID,
		Version: APIVersion,
		ReqData: reqDataCipherB64,
		Sign:    sign,
	}

	resp, err := c.post(ctx, msg)
	if err != nil {
		return nil, err
	}

	rawResp := resp.Body()

	if c.debug {
		log.Printf("[bosskg-debug] <<< Response Status: %d", resp.StatusCode())
		log.Printf("[bosskg-debug] <<< Response Headers: %v", resp.Header())
		log.Printf("[bosskg-debug] <<< Response Body (raw): %s", string(rawResp))
	}

	var respMsg RespMessage
	if err := json.Unmarshal(rawResp, &respMsg); err != nil {
		return nil, fmt.Errorf("unmarshal resp message: %w", err)
	}

	var plain []byte
	if respMsg.ResData != "" {
		if respMsg.Sign == "" {
			return nil, errors.New("missing resp signature")
		}
		if err := VerifyRSASHA1Base64([]byte(respMsg.ResData), respMsg.Sign, c.youfuPub); err != nil {
			return nil, fmt.Errorf("verify resp signature: %w", err)
		}

		plain, err = DecryptDESBase64(respMsg.ResData, c.desKey)
		if err != nil {
			return nil, fmt.Errorf("decrypt resData: %w", err)
		}
		respMsg.ResData = string(plain)

		if c.debug {
			log.Printf("[bosskg-debug] <<< ResData (decrypted): %s", string(plain))
		}
	}

	if respMsg.ResCode != "" && respMsg.ResCode != "0000" {
		return &respMsg, APIError{
			FunCode: respMsg.FunCode,
			Code:    respMsg.ResCode,
			Message: respMsg.ResMsg,
			Raw:     respMsg.ResData, // decrypted when ResData is present
		}
	}

	if out != nil {
		if len(plain) > 0 {
			if err := json.Unmarshal(plain, out); err != nil {
				return &respMsg, fmt.Errorf("unmarshal resData: %w", err)
			}
		}
	}

	return &respMsg, nil
}

func (c *Client) post(ctx context.Context, msg *ReqMessage) (*resty.Response, error) {
	b, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	headers := http.Header{
		"Content-Type": {"application/json;charset=UTF-8"},
	}
	if c.userAgent != "" {
		headers.Set("User-Agent", c.userAgent)
	}

	if c.debug {
		log.Printf("[bosskg-debug] >>> URL: POST %s", c.baseURL)
		log.Printf("[bosskg-debug] >>> Request Headers: %v", headers)
		log.Printf("[bosskg-debug] >>> Request Body: %s", string(b))
	}

	r := c.restyClient.R().
		SetContext(ctx).
		SetHeaderMultiValues(headers).
		SetBody(b)

	resp, err := r.Post(c.baseURL)
	if err != nil {
		return nil, fmt.Errorf("http post: %w", err)
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("http status=%d body=%s", resp.StatusCode(), string(resp.Body()))
	}
	return resp, nil
}
