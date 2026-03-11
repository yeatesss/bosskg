package bosskg

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"errors"
	"fmt"
)

func DeriveDESKey(desKey string) ([]byte, error) {
	if len(desKey) < 8 {
		return nil, errors.New("desKey must be at least 8 chars")
	}
	return []byte(desKey[:8]), nil
}

func EncryptDESBase64(plain []byte, key []byte) (string, error) {
	ct, err := encryptDESECB(plain, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ct), nil
}

func DecryptDESBase64(cipherBase64 string, key []byte) ([]byte, error) {
	ct, err := base64.StdEncoding.DecodeString(cipherBase64)
	if err != nil {
		return nil, fmt.Errorf("decode base64: %w", err)
	}
	return decryptDESECB(ct, key)
}

func encryptDESECB(plain []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	plain = pkcs5Pad(plain, bs)
	if len(plain)%bs != 0 {
		return nil, errors.New("plain is not full blocks after padding")
	}

	out := make([]byte, len(plain))
	for i := 0; i < len(plain); i += bs {
		block.Encrypt(out[i:i+bs], plain[i:i+bs])
	}
	return out, nil
}

func decryptDESECB(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	if len(ciphertext)%bs != 0 {
		return nil, errors.New("ciphertext is not full blocks")
	}

	out := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += bs {
		block.Decrypt(out[i:i+bs], ciphertext[i:i+bs])
	}
	return pkcs5Unpad(out)
}

func pkcs5Pad(b []byte, blockSize int) []byte {
	padding := blockSize - len(b)%blockSize
	return append(b, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func pkcs5Unpad(b []byte) ([]byte, error) {
	if len(b) == 0 {
		return nil, errors.New("empty")
	}
	n := int(b[len(b)-1])
	if n <= 0 || n > len(b) {
		return nil, errors.New("bad padding")
	}
	for i := len(b) - n; i < len(b); i++ {
		if int(b[i]) != n {
			return nil, errors.New("bad padding")
		}
	}
	return b[:len(b)-n], nil
}
