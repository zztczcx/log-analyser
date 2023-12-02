package httplog

import (
	"errors"
	"log"
	"regexp"
	"sync"
)

const (
        regexPattern = `^(\d+\.\d+\.\d+\.\d+) .* "(GET|POST|PUT|DELETE) ([^"]+) HTTP.*" \d+ \d+ ".*" ".*"`

)

func Parser(dataSource <-chan string, resultChan chan<-  Result, wg *sync.WaitGroup) {
        analyseChan := make(chan Record)
        go analyse(analyseChan, resultChan, wg)

	for d := range dataSource {
                record, err := parseData(d)
                if err != nil {
                        log.Println(err)
                }else{
                        analyseChan <- record
                }
	}
        close(analyseChan)
}

func parseData(line string) (Record, error) {
	re := regexp.MustCompile(regexPattern)
	matches := re.FindStringSubmatch(line)

	if len(matches) >= 4 {
		ip := matches[1]
		url := matches[3]
                return Record{ip, url}, nil

	} else {
                return Record{}, errors.New("Error parsing")
	}
}

func analyse(r <-chan Record, resultChan chan<- Result, wg *sync.WaitGroup){
        defer wg.Done()

        result := Result{
		VisitedUrls: make(map[string]int),
		ActiveIps:   make(map[string]int),
		UniqIps:     make(map[string]bool),
	}

        for record := range r {
                result.VisitedUrls[record.Url]++
                result.ActiveIps[record.Ip]++
                result.UniqIps[record.Ip] = true
        }

        resultChan <- result
}
