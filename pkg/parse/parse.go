package parse

import (
	"io"
	"strings"

	lgerrors "github.com/hamdanjaveed/lazygit/pkg/errors"
)

// ParseTable takes in table-formatted text and returns it in a structured format with whitespace trimmed out.
//
// It assumes the first row is a header and uses it to determine how wide each column is. Using this information
// it snips each subsequent row's columns and trims the cells to get just the data in each cell.
func ParseTable(t io.Reader) ([][]string, error) {
	b, err := io.ReadAll(t)
	if err != nil {
		return nil, lgerrors.Wrapf(err, "failed to read table")
	}
	rows := strings.Split(string(b), "\n")
	if len(rows) == 0 {
		return nil, lgerrors.Errorf("no data to parse")
	}

	widths, err := colWidthsFromHeaders(rows[0])
	if err != nil {
		return nil, lgerrors.Wrapf(err, "unable to parse col widths")
	}

	var data [][]string
	for i := 1; i < len(rows); i++ {
		row := make([]string, len(widths))
		r := rows[i]
		if len(r) < len(rows[0]) {
			r += strings.Repeat(" ", len(rows[0])-len(r))
		}
		if len(r) != len(rows[0]) {
			panic("wtf")
		}
		ptr := 0
		for j := 0; j < len(widths); j++ {
			end := ptr + widths[j]
			row[j] = strings.TrimSpace(r[ptr:end])
			ptr = end
		}
		data = append(data, row)
	}

	return data, nil
}

func colWidthsFromHeaders(headers string) ([]int, error) {
	if len(headers) == 0 {
		return nil, lgerrors.Errorf("headers is empty")
	}
	if headers[0] == ' ' {
		return nil, lgerrors.Errorf("was expecting first char of headers to be alphanumeric: '%s'", headers)
	}

	var widths []int
	currentColWidth := 0
	inWord := true
	for i, r := range headers {
		if i == len(headers)-1 {
			// We've reached the end of the headers row, save the current width and exit
			widths = append(widths, currentColWidth+1)
			break
		}

		if inWord {
			if r == ' ' {
				// We reached the end of the header text, but keep counting the column width
				inWord = false
			} else {
				// Continue the word
			}
		} else {
			if r == ' ' {
				// Continue the column
			} else {
				// We've reached the next column's header, save the current width and reset
				if currentColWidth == 0 {
					return nil, lgerrors.Errorf("unexpected 0 len col in headers: '%s'", headers)
				}

				widths = append(widths, currentColWidth)
				currentColWidth = 1
				inWord = true
				continue
			}
		}

		currentColWidth++
	}

	return widths, nil
}
