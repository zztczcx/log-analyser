package httplog

import (
	"log-analyser/stats"
)

type Statistic struct {
	UnitIpNum       int
	MostVisitedUrls []string
	MostActiveIps   []string
}

func (a *analyser) Statisticize(result Result) Statistic {
	s := Statistic{}
	numUniqIps := len(result.UniqIps)

	s.UnitIpNum = numUniqIps
	s.MostVisitedUrls = stats.TopMost(result.VisitedUrls, a.mostVisitedUrlsCount)
	s.MostActiveIps = stats.TopMost(result.ActiveIps, a.mostActiveIpsCount)

	return s
}
