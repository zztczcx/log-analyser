package httplog

import (
	"reflect"
	"testing"
)

func Test_Reducer(t *testing.T) {
        resultCh := make(chan Result)
        result1 := Result{
                VisitedUrls: map[string]int{
                        "http://www.google.com": 1,
                        "http://www.yahoo.com":  2,
                },
                ActiveIps: map[string]int{
                        "192.168.1.1": 1,
                        "192.168.1.2": 2,
                },
                UniqIps: map[string]bool{
                        "192.168.1.1": true,
                        "192.168.1.2": true,
                },
        }
        result2 := Result{
                VisitedUrls: map[string]int{
                        "http://www.google.com": 1,
                        "http://www.yahoo.com":  5,
                        "http://www.facebook.com":  5,
                },
                ActiveIps: map[string]int{
                        "192.168.1.1": 1,
                        "192.168.1.3": 2,
                },
                UniqIps: map[string]bool{
                        "192.168.1.1": true,
                        "192.168.1.3": true,
                },
        }

        go func() {
                resultCh <- result1
                resultCh <- result2
                close(resultCh)
        }()

        finalResult := Reducer(resultCh)

        if len(finalResult.UniqIps) != 3 {
                t.Errorf("Expected 3 unique ips")
        }

        expectedActiveIps := map[string]int{
                "192.168.1.1": 2,
                "192.168.1.2": 2,
                "192.168.1.3": 2,
        }

        if reflect.DeepEqual(finalResult.ActiveIps, expectedActiveIps) != true {
                t.Errorf("Expected %v\n got %v\n", expectedActiveIps, finalResult.ActiveIps)
        }

        expectedVistedUrls := map[string]int{
                "http://www.google.com": 2,
                "http://www.yahoo.com":  7,
                "http://www.facebook.com":  5,
        }

        if reflect.DeepEqual(finalResult.VisitedUrls, expectedVistedUrls) != true {
                t.Errorf("Expected %v\n got %v\n", expectedVistedUrls, finalResult.VisitedUrls)
        }
}
