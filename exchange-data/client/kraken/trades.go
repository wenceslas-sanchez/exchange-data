package kraken

import (
	"exchanges-data/exchange-data/client"
	"exchanges-data/exchange-data/utils"
	"time"
)

type Pair struct {
	Price, Volume             string
	Time                      float64
	Side, Kind, Miscellaneous string
}

type TradesData struct {
	Last  string
	Pairs []Pair
}

// TODO manage the output in the interface => genralize it (like (*Data, error))
func (k *KrakenClient) GetTrades(pair string, start, end time.Time) (any, error) {
	res, err := utils.Request(k.Client, client.GET, k.URL, nil, k.Header)
	if err != nil {
		return nil, err
	}
	var tradesData *TradesData
	err = utils.DecodeResponseJSON(res, tradesData)
	if err != nil {
		return nil, err
	}

	return tradesData, nil
}
