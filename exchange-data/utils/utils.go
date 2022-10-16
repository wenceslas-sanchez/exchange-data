package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// A Schema is the value used in json.NewDecoder(...).Decode().
type Schema interface{}

// Request send an HTTP request whatever the method (GET, POST, etc.).
//
// A client is a Go's HTTP client.
// A method is all HTTP available methods.
// An url is the address to send the request.
// A body is the response body. This argument is not mandatory. You can set its value to nil.
// A header is the request header. This argument is not mandatory. You can set its value to nil.
func Request(client *http.Client, method string, url string, body io.Reader, header *http.Header) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("connection error: %s", err)
	}
	if header != nil {
		req.Header = *header
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("bas response: %s", err)
	}

	// Maybe manage other than 200 status code than can be acceptable (203, etc.)
	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		bmsg, err := io.ReadAll(res.Body)
		msg := string(bmsg)
		if err != nil {
			msg = ""
			log.Printf("can't parse response body: %q", err)
		}
		return nil, fmt.Errorf("response error (error %q): %v", res.StatusCode, msg)
	}

	return res, nil
}

func DecodeResponseJSON[T any](res *http.Response, schema *T) error {
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(schema); err != nil {
		return err
	}

	return nil
}
