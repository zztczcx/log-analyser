package stats

import (
	"log-analyser/httplog"
	"slices"
	"testing"
)

func Test_Statisticize(t *testing.T) {
        result := httplog.Result{
                UniqIps: map[string]bool{
                        "192.168.1.1" : true,
                        "192.168.1.2" : true,
                        "192.168.1.3" : true,
                        "192.168.1.4" : true,
                },
                ActiveIps: map[string]int{
                        "192.168.1.1" : 2,
                        "192.168.1.2" : 3,
                        "192.168.1.3" : 1,
                        "192.168.1.4" : 5,
                },
                VisitedUrls: map[string]int{
                        "/abc" : 4,
                        "/bcd" : 4,
                        "/def" : 5,
                        "/efg" : 8,
                },
        }

        s := Statisticize(result)

        if s.UnitIpNum != 4 {
                t.Errorf("Statisticize UnitIpNum failed")
        }

        expectedMostActiveIps := []string{"192.168.1.4", "192.168.1.2", "192.168.1.1"}
        expectedMostVisitedUrls:= []string{"/efg", "/def", "/abc"}

        if slices.Equal(s.MostActiveIps, expectedMostActiveIps) != true {
                t.Errorf("Statisticize MostActiveIps failed")
        }
        if slices.Equal(s.MostVisitedUrls, expectedMostVisitedUrls) != true {
                t.Errorf("Statisticize MostVisitedUrls failed")
        }
}
