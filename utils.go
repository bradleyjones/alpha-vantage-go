package alphavantage

import "fmt"

func makeGetParams(params map[string]string) string {
	var output string
	for key, value := range params {
		output += fmt.Sprintf("&%s=%s", key, value)
	}
	return output
}
