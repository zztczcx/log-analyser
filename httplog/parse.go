package httplog

import (
	"errors"
	"log"
	"regexp"
	"sync"
)

const (
        regexPattern = `^(\d+\.\d+\.\d+\.\d+) .* "(GET|POST|PUT|DELETE) ([^"]+) HTTP.*" \d+ \d+ ".*" ".*"`
	parserCount = 5

)

func startParser(dataSource <-chan string, resultChan chan<- Result) {
	var wg sync.WaitGroup
	wg.Add(parserCount)

	for i := 0; i < parserCount; i++ {
		go Parse(dataSource, resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()
}

func Parse(dataSource <-chan string, resultChan chan<-  Result, wg *sync.WaitGroup) {
        collectChan := make(chan Record)
        go collect(collectChan, resultChan, wg)

	for d := range dataSource {
                record, err := parseData(d)
                if err != nil {
                        log.Println(err)
                }else{
                        collectChan <- record
                }
	}
        close(collectChan)
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

func collect(r <-chan Record, resultChan chan<- Result, wg *sync.WaitGroup){
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
