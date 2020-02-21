package alphavantage

import "testing"

func TestMakeGetParams(t *testing.T) {
	var testCases = map[string]struct {
		input    map[string]string
		expected string
	}{
		"Test Valid Symbol":        {map[string]string{"Symbol": "TEST"}, "&Symbol=TEST"},
		"Test nil params passed":   {nil, ""},
		"Test empty params passed": {map[string]string{}, ""},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			if output := makeGetParams(test.input); output != test.expected {
				t.Errorf("Test makeGetParams failed, inputted: %s, expected: %s, recieved: %s", test.input, test.expected, output)
			}
		})
	}
}
