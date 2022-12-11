package crypto

import (
	"fmt"

	gecko "github.com/superoo7/go-gecko/v3"
)

// Client holds data for gecko sdk.
type Client struct {
	cfg    Config
	client *gecko.Client
}

// New generates a new client.
func New(cfg Config) *Client {
	// create a new gecko client
	cg := gecko.NewClient(nil)

	return &Client{
		cfg:    cfg,
		client: cg,
	}
}

// Get returns the prize of crypto.
func (c *Client) Get() (string, float32, error) {
	price, err := c.client.SimpleSinglePrice(c.cfg.Crypto, c.cfg.Value)
	if err != nil {
		return "", 0, fmt.Errorf("get prize failed: %v", err)
	}

	return price.Currency, price.MarketPrice, nil
}
