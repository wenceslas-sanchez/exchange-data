package client

import "time"

type TradesData interface{}

type MarketClient interface {
	GetTrades(pair string, start, end time.Time) (any, error)
}
