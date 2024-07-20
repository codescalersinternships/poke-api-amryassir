package pkg

import (
	"net/http"
)

// Client - represents a client for making HTTP requests
type Client struct {
	httpClient *http.Client
	config     Config
}

// NewClient - initializes and returns a new instance of the Client struct
func NewClient(config Config) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
		config: config,
	}
}
