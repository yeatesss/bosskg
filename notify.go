package bosskg

import (
	"encoding/json"
	"fmt"
)

type NotifyMessage struct {
	FunCode FunCode `json:"funCode,omitempty"` // 接口编码
	MerID   string  `json:"merId,omitempty"`   // 商户号
	Version string  `json:"version,omitempty"` // 接口版本号
	ResData string  `json:"resData,omitempty"` // 通知业务数据（DES-ECB + Base64）
	ResCode string  `json:"resCode,omitempty"` // 响应码
	ResMsg  string  `json:"resMsg,omitempty"`  // 响应信息
	Sign    string  `json:"sign,omitempty"`    // 对 resData 的签名（RSA-SHA1 + Base64）
}

func (c *Client) ParseNotify(body []byte) (*NotifyMessage, []byte, error) {
	var n NotifyMessage
	if err := json.Unmarshal(body, &n); err != nil {
		return nil, nil, fmt.Errorf("unmarshal notify: %w", err)
	}
	if err := VerifyRSASHA1Base64([]byte(n.ResData), n.Sign, c.youfuPub); err != nil {
		return &n, nil, fmt.Errorf("verify notify signature: %w", err)
	}
	plain, err := DecryptDESBase64(n.ResData, c.desKey)
	if err != nil {
		return &n, nil, fmt.Errorf("decrypt notify resData: %w", err)
	}
	return &n, plain, nil
}
