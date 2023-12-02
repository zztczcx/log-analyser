package httplog

import (
	"errors"
	"log"
	"regexp"
)

const (
        regexPattern = `^(\S+) \S+ \S+ \[.*\] "(GET|POST|PUT|DELETE|HEAD) (\S+) .*" \d+ \d+$`

)

func Parser(dataSource <-chan string, resultChan chan<-  Result) {
        analyseChan := make(chan Record)
        go analyse(analyseChan, resultChan)

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

func analyse(r <-chan Record, resultChan chan<- Result){
        result := Result{}
        for record := range r {
                result.VisitedUrls[record.Url]++
                result.ActiveIps[record.Ip]++
                result.UniqIps[record.Ip] = true
        }

        resultChan <- result
}

