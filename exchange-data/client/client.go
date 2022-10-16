package client

import "time"

type MarketClient interface {
	GetTrades(pair string, start, end time.Time)
}
