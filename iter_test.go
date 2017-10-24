package iter

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestIter(t *testing.T) {
	rows := [][]string{}
	ch, done := make(chan []string), make(chan bool)

	go func() {
		for row := range ch {
			rows = append(rows, row)
		}
		done <- true
	}()

	file := filepath.Join("test", "file1.csv")
	err := Iter(file, true, ch)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	close(ch)
	<-done

	if len(rows) != 5 {
		t.Error("Expected 5, got", len(rows))
	}

	for i, row := range rows {
		id := row[0]
		expected := fmt.Sprintf("%v", i+1)
		if id != expected {
			t.Error("Expected", expected, "Got", id)
		}
	}
}

func TestDelimiter(t *testing.T) {
	rows := [][]string{}
	ch, done := make(chan []string), make(chan bool)

	go func() {
		for row := range ch {
			rows = append(rows, row)
		}
		done <- true
	}()

	file := filepath.Join("test", "file2.csv")
	err := Iter(file, true, ch, '|')
	if err != nil {
		t.Error("Unexpected error", err)
	}

	close(ch)
	<-done

	if len(rows) != 4 {
		t.Error("Expected 4, got", len(rows))
	}

	for i, row := range rows {
		id := row[0]
		expected := fmt.Sprintf("%v", i+1)
		if id != expected {
			t.Error("Expected", expected, "Got", id)
		}
	}
}

func TestNoHeader(t *testing.T) {
	rows := [][]string{}
	ch, done := make(chan []string), make(chan bool)

	go func() {
		for row := range ch {
			rows = append(rows, row)
		}
		done <- true
	}()

	file := filepath.Join("test", "file3.csv")
	err := Iter(file, false, ch)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	close(ch)
	<-done

	if len(rows) != 10 {
		t.Error("Expected 10, got", len(rows))
	}

	for i, row := range rows {
		id := row[0]
		expected := fmt.Sprintf("%v", i+1)
		if id != expected {
			t.Error("Expected", expected, "Got", id)
		}
	}
}
