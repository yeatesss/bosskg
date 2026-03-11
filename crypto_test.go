package bosskg

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"testing"
)

func TestParseKeys_Base64DER(t *testing.T) {
	t.Parallel()

	pri, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatal(err)
	}
	priDER, err := x509.MarshalPKCS8PrivateKey(pri)
	if err != nil {
		t.Fatal(err)
	}
	pubDER, err := x509.MarshalPKIXPublicKey(&pri.PublicKey)
	if err != nil {
		t.Fatal(err)
	}

	gotPri, err := ParsePrivateKey(base64.StdEncoding.EncodeToString(priDER))
	if err != nil {
		t.Fatal(err)
	}
	gotPub, err := ParsePublicKey(base64.StdEncoding.EncodeToString(pubDER))
	if err != nil {
		t.Fatal(err)
	}

	sig, err := SignRSASHA1Base64([]byte("hello"), gotPri)
	if err != nil {
		t.Fatal(err)
	}
	if err := VerifyRSASHA1Base64([]byte("hello"), sig, gotPub); err != nil {
		t.Fatal(err)
	}
}
