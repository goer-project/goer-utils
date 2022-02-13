package aes

func NewClient(key string, iv string) *Client {
	return &Client{
		Key: key,
		Iv:  iv,
	}
}

type Client struct {
	Key string
	Iv  string
}
