package aes

import (
	"crypto/aes"
	"crypto/cipher"
	b64 "encoding/base64"
)

func (client *Client) CBCEncrypt(text string) string {
	block, err := aes.NewCipher([]byte(client.Key))
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()
	byteIv := []byte(client.Iv)
	iv := byteIv[:blockSize]

	src := PKCS7Padding([]byte(text), blockSize)
	dst := make([]byte, len(src))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(dst, src)

	return b64.StdEncoding.EncodeToString(dst)
}

func (client *Client) CBCDecrypt(encrypted string) string {
	block, err := aes.NewCipher([]byte(client.Key))
	if err != nil {
		panic(err)
	}

	src, _ := b64.StdEncoding.DecodeString(encrypted)

	blockSize := block.BlockSize()
	byteIv := []byte(client.Iv)
	iv := byteIv[:blockSize]

	dst := make([]byte, len(src))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(dst, src)

	dst, err = PKCS7UnPadding(dst)
	if err != nil {
		panic(err)
	}

	return string(dst)
}
