package bosskg

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

type PrivateKey = rsa.PrivateKey
type PublicKey = rsa.PublicKey

func ParsePrivateKey(s string) (*PrivateKey, error) {
	der, err := parseKeyBytes(s)
	if err != nil {
		return nil, err
	}

	if pk, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		rsaKey, ok := pk.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("pkcs8 private key is not RSA")
		}
		return rsaKey, nil
	}
	if rsaKey, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return rsaKey, nil
	}
	return nil, errors.New("unsupported private key format (expect PKCS8 or PKCS1)")
}

func ParsePublicKey(s string) (*PublicKey, error) {
	der, err := parseKeyBytes(s)
	if err != nil {
		return nil, err
	}
	pk, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return nil, fmt.Errorf("parse pkix public key: %w", err)
	}
	rsaKey, ok := pk.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("public key is not RSA")
	}
	return rsaKey, nil
}

func parseKeyBytes(s string) ([]byte, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, errors.New("empty key")
	}

	if strings.Contains(s, "BEGIN ") {
		block, _ := pem.Decode([]byte(s))
		if block == nil || len(block.Bytes) == 0 {
			return nil, errors.New("invalid pem key")
		}
		return block.Bytes, nil
	}

	s = stripAllWhitespace(s)
	der, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, errors.New("invalid base64 key")
	}
	return der, nil
}

func SignRSASHA1Base64(data []byte, privateKey *PrivateKey) (string, error) {
	h := sha1.Sum(data)
	sig, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, h[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sig), nil
}

func VerifyRSASHA1Base64(data []byte, signatureBase64 string, publicKey *PublicKey) error {
	sig, err := base64.StdEncoding.DecodeString(strings.TrimSpace(signatureBase64))
	if err != nil {
		return fmt.Errorf("decode signature base64: %w", err)
	}
	h := sha1.Sum(data)
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA1, h[:], sig)
}

func stripAllWhitespace(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch r {
		case ' ', '\n', '\r', '\t':
			continue
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}
