package twitter

type trend struct {
	Name string
}

type trendWrapper struct {
	Trends []trend
}

type trendResponse []trendWrapper
