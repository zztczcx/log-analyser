package httplog

func Reducer(resultChan <-chan Result) Result{
        finalResult := Result{}
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
