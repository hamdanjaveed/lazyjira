package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColWidthsFromHeaders(t *testing.T) {
	tests := []struct {
		name     string
		headers  string
		expected []int
		err      string
	}{
		{
			name:     "valid - simple headers (simulated)",
			headers:  "10        5    8       ",
			expected: []int{10, 5, 8},
		},
		{
			name:     "valid - simple headers (real)",
			headers:  "TYPE            KEY     SUMMARY                                                                                                                                                                                                                                                 STATUS          ASSIGNEE        REPORTER        PRIORITY        RESOLUTION      CREATED                  UPDATED                  LABELS",
			expected: []int{16, 8, 248, 16, 16, 16, 16, 16, 25, 25, 6},
		},
		{
			name:     "valid - single column",
			headers:  "TYPE",
			expected: []int{4},
		},
		{
			name:     "valid - whitespace in last column",
			headers:  "TYPE            KEY     ",
			expected: []int{16, 8},
		},
		{
			name:    "invalid - no data",
			headers: "",
			err:     "headers is empty",
		},
		{
			name:    "invalid - whitespace at the beginning",
			headers: " TYPE",
			err:     "first char of headers to be alphanumeric",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			widths, err := colWidthsFromHeaders(test.headers)
			if test.err != "" {
				assert.ErrorContains(t, err, test.err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expected, widths)
		})
	}
}
