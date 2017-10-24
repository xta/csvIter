package iter

import (
	"encoding/csv"
	"io"
	"os"
)

const defaultDelimiter = ','

// Iter iterates over a CSV file's rows to return the rows without the header.
func Iter(file string, hasHeader bool, ch chan []string, delimiter ...rune) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	c := csv.NewReader(f)

	if len(delimiter) > 0 {
		c.Comma = delimiter[0]
	} else {
		c.Comma = defaultDelimiter
	}

	if hasHeader {
		_, err := c.Read()
		if err != nil {
			return err
		}
	}

	for {
		row, err := c.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		ch <- row
	}

	return nil
}
