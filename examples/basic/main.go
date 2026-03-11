package main

import (
	"context"
	"log"
	"os"

	"github.com/yeatesss/bosskg"
)

func main() {
	cli, err := bosskg.NewClient(bosskg.Config{
		BaseURL:            env("BOSSKG_BASE_URL", bosskg.TestGatewayURL),
		MerchantID:         mustEnv("BOSSKG_MER_ID"),
		DESKey:             mustEnv("BOSSKG_DES_KEY"),
		MerchantPrivateKey: mustEnv("BOSSKG_MER_PRI_KEY"),
		YoufuPublicKey:     mustEnv("BOSSKG_YOUFU_PUB_KEY"),
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

func mustEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("missing env %s", k)
	}
	return v
}

func env(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}
