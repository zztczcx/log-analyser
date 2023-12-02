package httplog

import (
	"sync"
	"testing"
)

func Test_parseData(t *testing.T) {
        line := `177.71.128.21 - - [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`

        record, _ := parseData(line)

        if record.Ip != "177.71.128.21" {
                t.Errorf("parse Ip failed")
        }

        if record.Url != "/intranet-analytics/" {
                t.Errorf("parse url failed")
        }
}

func Test_analyse(t *testing.T) {
        recordCh := make(chan Record)
        resultCh := make(chan Result)
        var wg sync.WaitGroup
        wg.Add(2)

        go func() {
                recordCh <- Record{Ip: "177.71.182.21", Url: "/abc"}
                recordCh <- Record{Ip: "177.71.182.21", Url: "/abc"}
                recordCh <- Record{Ip: "177.71.182.22", Url: "/def"}
                close(recordCh)
        }()
        go analyse(recordCh, resultCh, &wg)

        r := <-resultCh


        if len(r.UniqIps) != 2 {
                t.Errorf("analyse UniqIps failed")
        }


        if r.ActiveIps["177.71.182.21"] != 2 || r.ActiveIps["177.71.182.22"]  != 1{
                 t.Errorf("analyse ActiveIps failed")
        }

        if r.VisitedUrls["/abc"] != 2 || r.VisitedUrls["/def"]  != 1{
                 t.Errorf("analyse VisitedUrls failed")
        }
}
