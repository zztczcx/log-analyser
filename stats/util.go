package stats

import (
	"log-analyser/httplog"
	"sort"
)

type Statistic struct {
	UnitIpNum       int
	MostVisitedUrls []string
	MostActiveIps   []string
}

func Statisticize(result httplog.Result) Statistic {
	s := Statistic{}
	numUniqIps := len(result.UniqIps)

	s.UnitIpNum = numUniqIps
	s.MostVisitedUrls = getFirstNKeys(getKeysBySortMapByValue(result.VisitedUrls), 3)
	s.MostActiveIps = getFirstNKeys(getKeysBySortMapByValue(result.ActiveIps), 3)

	return s
}

func getKeysBySortMapByValue(m map[string]int) *[]string {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})

        return &keys
}

func getFirstNKeys(keys *[]string, n int) []string{
        if n > len(*keys) {
                return *keys
        }
        return (*keys)[:n]
}
