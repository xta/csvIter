# CSV Iter

CSV Iter is a wrapper around golang's *encoding/csv* that provides access to the CSV row contents via a channel.

## Usage

`main.go`

	package main

	import (
		"fmt"
		csv "github.com/xta/csvIter"
	)

	func main() {
		ch, done := make(chan []string), make(chan bool)

		go func() {
			for row := range ch {
				// do something with the row `[]string`
			}
			done <- true
		}()

		if err := csv.Iter("path/to/file.csv", true, ch); err != nil {
			fmt.Println("Error", err)
		}

		close(ch)
		<-done
	}

## Tests

	make
