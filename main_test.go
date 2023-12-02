package main

import (
	"os"
	"strings"
	"testing"
)

const (
        ExpectedOutput = `The number of unique IP addresses: 11
The top 3 most visited URLs: [/docs/manage-websites/ / /asset.css]
The top 3 most active IP addresses: [168.41.191.40 177.71.128.21 50.112.00.11]`
)

func TestMain(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	out := make([]byte, 1024)
	r.Read(out)
        got := string(out)

	if strings.Compare(ExpectedOutput, got) == 0 {
		t.Errorf("\nExpected: %s\nGot: %s\n", ExpectedOutput, got)
	}
}
