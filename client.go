package alphavantage

type AVClient struct {
	apiKey string
}

// Return a new alpha vantage client
func NewClient(apiKey string) *AVClient {
	return &AVClient{
		apiKey: apiKey,
	}
}
