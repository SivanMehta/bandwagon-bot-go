package twitter

type trend struct {
	Name         string
	Tweet_volume int
}

type trendWrapper struct {
	Trends []trend
}

type trendResponse []trendWrapper

type tweet struct {
	Text string
}

type metadata struct {
	Next_results string
}

type tweetsResponse struct {
	Statuses        []tweet
	Search_metadata metadata
}
