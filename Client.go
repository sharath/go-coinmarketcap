package go_coinmarketcap

import (
	"net/http"
	"time"
)

const (
	baseURL        = "https://pro-api.coinmarketcap.com/"
	apiURL         = "https://pro-api.coinmarketcap.com/v1/"
	cryptocurrency = "cryptocurrency/"
)

type Client struct {
	apiKey string
	client *http.Client
}

func NewClient(apiKey string) *Client {
	if apiKey == "" {
		return nil
	}
	c := new(Client)
	c.apiKey = apiKey
	c.client = &http.Client{Timeout: 10 * time.Second}
	return c
}
