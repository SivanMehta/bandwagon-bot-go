package twitter

import (
	"encoding/json"
)

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

func extractTweetsfromResponse(response tweetsResponse) []string {
	data := response.Statuses
	tweets := make([]string, len(data))

	for i := range data {
		tweets[i] = data[i].Text
	}

	return tweets
}

func getNextTweets(nextURL string, n int, topic string) []string {
	if n < 0 {
		return []string{}
	}

	responseBytes := makeAuthedRequest("GET", "1.1/search/tweets.json"+nextURL)

	var parsed tweetsResponse
	json.Unmarshal(responseBytes, &parsed)
	tweets := extractTweetsfromResponse(parsed)
	if len(tweets) == 0 {
		return []string{}
	}

	nextPage := parsed.Search_metadata.Next_results
	return append(tweets, getNextTweets(nextPage, n-len(tweets), topic)...)
}
