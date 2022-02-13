package aes

import (
	"crypto/aes"
	"crypto/cipher"
	b64 "encoding/base64"
)

func (client *Client) CFBEncrypt(text string) string {
	block, err := aes.NewCipher([]byte(client.Key))
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()
	byteIv := []byte(client.Iv)
	iv := byteIv[:blockSize]

	src := []byte(text)
	dst := make([]byte, len(src))
	mode := cipher.NewCFBEncrypter(block, iv)
	mode.XORKeyStream(dst, src)

	return b64.StdEncoding.EncodeToString(dst)
}

func (client *Client) CFBDecrypt(encrypted string) string {
	block, err := aes.NewCipher([]byte(client.Key))
	if err != nil {
		panic(err)
	}

	src, _ := b64.StdEncoding.DecodeString(encrypted)

	blockSize := block.BlockSize()
	byteIv := []byte(client.Iv)
	iv := byteIv[:blockSize]

	dst := make([]byte, len(src))
	mode := cipher.NewCFBDecrypter(block, iv)
	mode.XORKeyStream(dst, src)

	return string(dst)
}
