package httplog

import (
	"bufio"
	"log"
	"os"
)

func produce(s string) <-chan string {
	dataSource := make(chan string)

	go func() {
		f, err := os.Open(s)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			dataSource <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		close(dataSource)

	}()

	return dataSource
}
