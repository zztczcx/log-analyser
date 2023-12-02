package main

import (
	"bufio"
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
	dataSource := producer("./programming-task-example-data.log")
        resultChan := make(chan httplog.Result)
        var wg sync.WaitGroup
        wg.Add(parserCount)

	for i := 0; i < parserCount; i++ {
		go httplog.Parser(dataSource, resultChan, &wg)
	}

        go func(){
                wg.Wait()
                close(resultChan)
        }()

        finalResult := httplog.Reducer(resultChan)
        report(stats.Statisticize(finalResult))

}


func report(s stats.Statistic){

        log.Printf("The number of unique IP addresses: %v\n", s.UnitIpNum)
        log.Printf("The top 3 most visited URLs: %v\n", s.MostVisitedUrls)
        log.Printf("The top 3 most active IP addresses: %v\n", s.MostActiveIps)
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

