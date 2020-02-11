package alphavantage

import (
	"encoding/json"
	"fmt"
)

type AVQuote struct {
	GlobalQuote Quote `json:"Global Quote"`
}

type Quote struct {
	Symbol           string `json:"01. symbol"`
	Open             string `json:"02. open"`
	High             string `json:"03. high"`
	Low              string `json:"04. low"`
	Price            string `json:"05. price"`
	Volume           string `json:"06. volume"`
	LatestTradingDay string `json:"07. latest trading day"`
	PreviousClose    string `json:"08. previous close"`
	Change           string `json:"09. change"`
	ChangePercent    string `json:"10. change percent"`
}

func parseQuote(resp []byte) (*Quote, error) {
	var quote AVQuote

	err := json.Unmarshal(resp, &quote)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling quote %w, response: %s", err, string(resp))
	}

	return &quote.GlobalQuote, nil
}
