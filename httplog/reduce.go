package httplog

func Reduce(resultChan <-chan Result) Result{
        finalResult := Result{
		VisitedUrls: make(map[string]int),
		ActiveIps:   make(map[string]int),
		UniqIps:     make(map[string]bool),
	}
        for result := range resultChan {
                for k, v := range result.VisitedUrls{
                        finalResult.VisitedUrls[k] += v
                }

                for k, v := range result.ActiveIps{
                        finalResult.ActiveIps[k] += v
                }

                for k, v := range result.UniqIps{
                        finalResult.UniqIps[k] = v
                }
        }

        return finalResult
}
