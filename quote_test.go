package alphavantage

import "testing"

const (
	valid_json_response = `{
    "Global Quote": {
        "05. price": "183.7299"
				}
		}`
	invalid_json_response = `{
    "Global Quote
        "05. price": "183.7299",
				}
		}`
)

func TestParseQuote(t *testing.T) {
	var testCases = map[string]struct {
		input        []byte
		expected     string
		expected_err bool
	}{
		"Test Valid JSON":   {[]byte(valid_json_response), "183.7299", false},
		"Test Invalid JSON": {[]byte(invalid_json_response), "", true},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			output, err := parseQuote(test.input)
			if err != nil && !test.expected_err {
				t.Errorf("Did not expect error but recieved: %s", err)
			} else if !test.expected_err && output.Price != test.expected {
				t.Errorf("Expected response: %s but recieved: %s", test.expected, output.Price)
			}
		})
	}
}
