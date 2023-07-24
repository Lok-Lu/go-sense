package sense

import (
	"net/http"
)

const (
	sensaAPIURLv1                  = "https://api.sensenova.cn/v1/llm"
	defaultEmptyMessagesLimit uint = 300
)

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	accessKey string
	secretKey string

	authToken          string
	BaseURL            string
	HTTPClient         *http.Client
	EmptyMessagesLimit uint
}

func DefaultConfig(ak, sk string) (ClientConfig, error) {
	authToken, err := EncodeJwtToken(ak, sk)
	if err != nil {
		return ClientConfig{}, err
	}

	return ClientConfig{
		accessKey:  ak,
		secretKey:  sk,
		BaseURL:    sensaAPIURLv1,
		authToken:  authToken,
		HTTPClient: &http.Client{},

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}, nil
}

func (c ClientConfig) WithHttpClientConfig(client *http.Client) ClientConfig {
	c.HTTPClient = client
	return c
}

func (ClientConfig) String() string {
	return "<OpenAI API ClientConfig>"
}
