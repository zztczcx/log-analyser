package main

import (
	"flag"
	"fmt"

	"log-analyser/httplog"
)

func main() {
	i := flag.String("input", "./testdata/http.log", "Source file")
	flag.Parse()

	if *i == "" {
		panic("Missing log file")
	}

        config := httplog.AnalyserConfig{
                MostVisitedUrlsCount: 3,
                MostActiveIpsCount:   3,
                LogFile:              i,
        }

	analyser := httplog.NewAnalyser(&config)
	statistic := analyser.Analyse()
	report(statistic, config)

}

func report(s httplog.Statistic, c httplog.AnalyserConfig) {
	fmt.Printf("The number of unique IP addresses: %v\n", s.UnitIpNum)
	fmt.Printf("The top %d most visited URLs: %v\n", c.MostVisitedUrlsCount, s.MostVisitedUrls)
	fmt.Printf("The top %d most active IP addresses: %v\n", c.MostActiveIpsCount, s.MostActiveIps)
}
