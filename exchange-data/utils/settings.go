package utils

import (
	"net/http"
	"os"
)

var KRAKEN_HEADER = http.Header{
	"API-Key": {os.Getenv("KRAKEN_API_KEY")},
}
