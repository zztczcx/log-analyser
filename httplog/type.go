package httplog

type Record struct {
	Ip  string
	Url string
}

type Result struct {
        VisitedUrls map[string]int
        ActiveIps map[string]int
        UniqIps map[string]bool
}

