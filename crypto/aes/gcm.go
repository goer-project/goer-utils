package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	b64 "encoding/base64"
	"io"
)

func RandBytes(length int) (data []byte, err error) {
	data = make([]byte, length)
	if _, err = io.ReadFull(rand.Reader, data); err != nil {
		return nil, err
	}
	return data, err
}

func (client *Client) GCMEncrypt(text string) string {
	block, err := aes.NewCipher([]byte(client.Key))
	if err != nil {
		panic(err)
	}

	mode, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce, err := RandBytes(mode.NonceSize())
	if err != nil {
		panic(err)
	}

	src := []byte(text)
	dst := mode.Seal(nonce, nonce, src, nil)

	return b64.StdEncoding.EncodeToString(dst)
}

func (client *Client) GCMDecrypt(encrypted string) string {
	block, err := aes.NewCipher([]byte(client.Key))
	if err != nil {
		panic(err)
	}

	mode, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	ciphertext, _ := b64.StdEncoding.DecodeString(encrypted)

	nonceSize := mode.NonceSize()
	if len(ciphertext) < nonceSize {
		panic("ciphertext too short")
	}

	var nonce []byte
	nonce, ciphertext = ciphertext[:nonceSize], ciphertext[nonceSize:]
	dst, err := mode.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	return string(dst)
}
