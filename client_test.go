package bosskg

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestClient_Do_QueryMerIDBalance(t *testing.T) {
	t.Parallel()

	merPri, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatal(err)
	}
	youfuPri, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatal(err)
	}

	merPriDER, err := x509.MarshalPKCS8PrivateKey(merPri)
	if err != nil {
		t.Fatal(err)
	}
	youfuPubDER, err := x509.MarshalPKIXPublicKey(&youfuPri.PublicKey)
	if err != nil {
		t.Fatal(err)
	}

	desKey := "ABCDEFGH"

	hc := &http.Client{
		Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
			reqBody, err := io.ReadAll(r.Body)
			if err != nil {
				return nil, err
			}

			var reqMsg ReqMessage
			if err := json.Unmarshal(reqBody, &reqMsg); err != nil {
				t.Fatalf("decode req: %v", err)
			}

			if err := VerifyRSASHA1Base64([]byte(reqMsg.ReqData), reqMsg.Sign, &merPri.PublicKey); err != nil {
				t.Fatalf("verify req sign: %v", err)
			}
			plain, err := DecryptDESBase64(reqMsg.ReqData, []byte(desKey))
			if err != nil {
				t.Fatalf("decrypt reqData: %v", err)
			}

			var apiReq QueryMerIDBalanceReq
			if err := json.Unmarshal(plain, &apiReq); err != nil {
				t.Fatalf("unmarshal api req: %v", err)
			}
			if apiReq.ProviderID != 30481 {
				t.Fatalf("providerID=%d", apiReq.ProviderID)
			}

			apiResp := QueryMerIDBalanceResp{
				Balance:    "174310397",
				ProviderID: 30481,
			}
			apiRespPlain, _ := json.Marshal(apiResp)
			resDataCipher, err := EncryptDESBase64(apiRespPlain, []byte(desKey))
			if err != nil {
				t.Fatalf("encrypt resp: %v", err)
			}
			sign, err := SignRSASHA1Base64([]byte(resDataCipher), youfuPri)
			if err != nil {
				t.Fatalf("sign resp: %v", err)
			}

			respMsg := RespMessage{
				ReqID:   reqMsg.ReqID,
				FunCode: reqMsg.FunCode,
				MerID:   reqMsg.MerID,
				Version: APIVersion,
				ResData: resDataCipher,
				ResCode: "0000",
				ResMsg:  "成功",
				Sign:    sign,
			}
			respBody, _ := json.Marshal(respMsg)

			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader(respBody)),
				Request:    r,
			}, nil
		}),
	}

	cli, err := NewClient(Config{
		BaseURL:            "http://example.invalid",
		MerchantID:         "m1",
		DESKey:             desKey,
		MerchantPrivateKey: base64.StdEncoding.EncodeToString(merPriDER),
		YoufuPublicKey:     base64.StdEncoding.EncodeToString(youfuPubDER),
	}, WithRestyClient(resty.NewWithClient(hc)), WithRequestIDFunc(func() string { return "rid" }))
	if err != nil {
		t.Fatal(err)
	}

	out, err := cli.QueryMerIDBalance(context.Background(), QueryMerIDBalanceReq{ProviderID: 30481})
	if err != nil {
		t.Fatal(err)
	}
	if out.ProviderID != 30481 || out.Balance != "174310397" {
		t.Fatalf("unexpected resp: %+v", out)
	}
}

func TestClient_Do_APIError(t *testing.T) {
	t.Parallel()

	merPri, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatal(err)
	}
	youfuPri, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatal(err)
	}

	merPriDER, _ := x509.MarshalPKCS8PrivateKey(merPri)
	youfuPubDER, _ := x509.MarshalPKIXPublicKey(&youfuPri.PublicKey)

	desKey := "ABCDEFGH"

	hc := &http.Client{
		Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
			reqBody, err := io.ReadAll(r.Body)
			if err != nil {
				return nil, err
			}
			var reqMsg ReqMessage
			_ = json.Unmarshal(reqBody, &reqMsg)

			resDataCipher, _ := EncryptDESBase64([]byte(`{"resCode":"9999","resMsg":"失败"}`), []byte(desKey))
			sign, _ := SignRSASHA1Base64([]byte(resDataCipher), youfuPri)
			respBody, _ := json.Marshal(RespMessage{
				ReqID:   reqMsg.ReqID,
				FunCode: reqMsg.FunCode,
				MerID:   reqMsg.MerID,
				Version: APIVersion,
				ResData: resDataCipher,
				ResCode: "9999",
				ResMsg:  "失败",
				Sign:    sign,
			})

			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader(respBody)),
				Request:    r,
			}, nil
		}),
	}

	cli, err := NewClient(Config{
		BaseURL:            "http://example.invalid",
		MerchantID:         "m1",
		DESKey:             desKey,
		MerchantPrivateKey: base64.StdEncoding.EncodeToString(merPriDER),
		YoufuPublicKey:     base64.StdEncoding.EncodeToString(youfuPubDER),
	}, WithRestyClient(resty.NewWithClient(hc)))
	if err != nil {
		t.Fatal(err)
	}

	_, err = cli.QueryMerIDBalance(context.Background(), QueryMerIDBalanceReq{ProviderID: 1})
	if err == nil {
		t.Fatal("expected error")
	}
	if _, ok := err.(APIError); !ok {
		t.Fatalf("want APIError, got %T: %v", err, err)
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}
