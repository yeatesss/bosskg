# bosskg (Go SDK)

基于官方/第三方 `bosskg-demo-master/golang` 样例封装的 Go SDK：把公共报文、DES-ECB(Base64) 加解密、RSA-SHA1(Base64) 签名验签和常用业务接口收敛到一个 `Client` 里。

## 安装

```bash
go get github.com/yeatesss/bosskg
```

## 快速开始

```go
package main

import (
	"context"
	"log"

	"github.com/yeatesss/bosskg"
)

func main() {
	cli, err := bosskg.NewClient(bosskg.Config{
		BaseURL:            bosskg.TestGatewayURL, // 或生产网关地址
		MerchantID:         "YOUR_MER_ID",
		DESKey:             "YOUR_DES_KEY",
		MerchantPrivateKey: "YOUR_RSA_PRIVATE_KEY",
		YoufuPublicKey:     "YOUFU_RSA_PUBLIC_KEY",
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := cli.QueryMerIDBalance(context.Background(), bosskg.QueryMerIDBalanceReq{
		ProviderID: 30481,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", resp)
}
```

说明：
- `DESKey` 默认按样例规则取前 8 个字符作为 DES Key（DES-ECB + PKCS5Padding）。
- RSA 公私钥支持直接传 PEM 文本，或传“无头无尾”的 Base64(DER) 字符串（与样例一致）。

## 回调验签解密

```go
notify, plain, err := cli.ParseNotify(bodyBytes)
_ = notify
_ = plain // plain 是解密后的 resData（JSON bytes）
_ = err
```
