package kraken

import (
	"net/http"
	"time"

	"exchanges-data/exchange-data/utils"
)

type KrakenResponse struct {
	Error error
	data  *utils.Schema
}

// A KrakenClient store all the needed information to request the Kraken's API.
type KrakenClient struct {
	// Go's base HTTP client.
	Client *http.Client
	// Kraken's API URL
	URL string `default:"https://api.kraken.com/0"`
	*http.Header
}

// New allow to construct a KrakenClient.
//
// client argument is not mandatory, if not provided, then the Go's HTTP client is used.
func New(url string, header *http.Header, client *http.Client) *KrakenClient {
	if client == nil {
		client = &http.Client{}
	}
	return &KrakenClient{client, url, header}
}

func (k *KrakenClient) GetTrades(pair string, start, end time.Time) {

}
