package twitter

import (
	"encoding/json"
	"net/url"
	"sort"
)

var accessToken = ""
var token = ""
var secret = ""
var currentTrends []string

//
// GetTweets will return a body of text from trending topics
// This is currently just the text of the Declaration of Independence, but stay tuned.
//
func GetTweets(trend string) []string {
	// exactly containing the trend, without tweets that contain links
	query := trend + " -filter:links"
	query = url.QueryEscape(query)
	responseBytes := makeAuthedRequest("GET", "1.1/search/tweets.json?q="+query+"&include_entities=false&lang=en&result_type=popular")

	var parsed tweetsResponse
	json.Unmarshal(responseBytes, &parsed)
	data := parsed.Statuses
	tweets := make([]string, len(data))

	for i := range data {
		tweets[i] = data[i].Text
	}

	return tweets
}

//
// FetchTrends goes to twitter and returns a list of the top 5 trending topics
// the woeid of 2352824 hard codes this to the US
//
func FetchTrends() []string {
	responseBytes := makeAuthedRequest("GET", "1.1/trends/place.json?id=2352824")

	var parsed trendResponse
	json.Unmarshal(responseBytes, &parsed)

	data := parsed[0].Trends
	sort.Slice(data, func(i, j int) bool {
		return data[i].Tweet_volume > data[j].Tweet_volume
	})

	// only take the top 5 trends so we can stay relevant
	n := 5
	trends := make([]string, n)

	for i := 0; i < n; i++ {
		trends[i] = data[i].Name
	}

	currentTrends = trends
	return trends
}

//
// GetTrends simply returns whatever is trending right now
//
func GetTrends() []string {
	return currentTrends
}
