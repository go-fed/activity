package pub

import (
	"testing"
	"github.com/go-test/deep"
)

func TestSuccessfulParsing(t *testing.T) {
	table := []struct{
		name string
		input string
		expected mediaTypeHeaderResult
	}{
		{
			name: "",
			input: "",
			expected: mediaTypeHeaderResult{},
		},
	}
	for _, test := range table {
	}
}
