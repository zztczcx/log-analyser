package httplog

type Analyser interface {
	Analyse() Statistic
}

type AnalyserConfig struct {
	MostVisitedUrlsCount int
	MostActiveIpsCount   int
        LogFile              *string
}

type analyser struct {
	logFile              *string
	mostVisitedUrlsCount int
	mostActiveIpsCount   int
}

func NewAnalyser(c *AnalyserConfig) Analyser {
	return &analyser{
		logFile:     c.LogFile,
                mostVisitedUrlsCount: c.MostVisitedUrlsCount,
                mostActiveIpsCount:   c.MostActiveIpsCount,
	}
}

func (a *analyser) Analyse() Statistic {
	dataSource := produce(*a.logFile)
	resultChan := make(chan Result)
	startParser(dataSource, resultChan)
	finalResult := Reduce(resultChan)
	return a.Statisticize(finalResult)
}
