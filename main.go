package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"log-analyser/httplog"
	"log-analyser/stats"
)

const (
	parserCount = 5

)

func main() {
        i := flag.String("input", "./testdata/http.log", "Source file")
        flag.Parse()

        if *i == "" {
                panic("Missing log file")
        }

	dataSource := producer(*i)
        resultChan := make(chan httplog.Result)
        startParser(dataSource, resultChan, parserCount)
        finalResult := httplog.Reducer(resultChan)

        report(stats.Statisticize(finalResult))

}

func startParser(dataSource <-chan string, resultChan chan<- httplog.Result, parserCount int) {
        var wg sync.WaitGroup
        wg.Add(parserCount)

	for i := 0; i < parserCount; i++ {
		go httplog.Parser(dataSource, resultChan, &wg)
	}

        go func(){
                wg.Wait()
                close(resultChan)
        }()
}


func report(s stats.Statistic){
        fmt.Printf("The number of unique IP addresses: %v\n", s.UnitIpNum)
        fmt.Printf("The top 3 most visited URLs: %v\n", s.MostVisitedUrls)
        fmt.Printf("The top 3 most active IP addresses: %v\n", s.MostActiveIps)
}

func producer(s string) <-chan string {
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

