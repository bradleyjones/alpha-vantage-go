package alphavantage

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseURL = "https://www.alphavantage.co/"
)

type Client struct {
	apiKey string
}

// Return a new alpha vantage client
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

func (c *Client) Request(function string, params map[string]string) ([]byte, error) {
	response, err := http.Get(fmt.Sprintf("%squery?function=%s%s&apikey=%s", baseURL, function, makeGetParams(params), c.apiKey))
	if err != nil {
		return nil, fmt.Errorf("Error in Get request %w", err)
	}

	// Check response codes
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Error with request, recieved status code: %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body %w", err)
	}
	response.Body.Close()
	return body, nil
}

func (c *Client) GetQuote(symbol string) (*Quote, error) {
	params := map[string]string{
		"symbol": symbol,
	}
	body, err := c.Request("GLOBAL_QUOTE", params)
	if err != nil {
		return nil, err
	}

	return parseQuote(body)
}
