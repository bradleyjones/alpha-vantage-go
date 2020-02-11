package alphavantage

type Client struct {
	apiKey string
}

// Return a new alpha vantage client
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

func (c *Client) GetAPIKey() string {
	return c.apiKey
}
