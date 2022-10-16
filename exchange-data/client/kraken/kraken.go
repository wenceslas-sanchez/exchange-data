package kraken

import (
	"net/http"
)

const KRAKEN_URL = "https://api.kraken.com/0"

type KrakenResponse struct {
	Error  error
	Result any
}

// A KrakenClient store all the needed information to request the Kraken's API.
type KrakenClient struct {
	// Go's base HTTP client.
	Client *http.Client
	*http.Header
	// Kraken's API URL
	URL string
}

// New allow to construct a KrakenClient.
//
// client argument is not mandatory, if not provided, then the Go's HTTP client is used.
func New(header *http.Header, client *http.Client, url string) *KrakenClient {
	if client == nil {
		client = &http.Client{}
	}
	if url == "" {
		url = KRAKEN_URL
	}
	return &KrakenClient{client, header, url}
}
