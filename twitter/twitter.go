package twitter

import (
	"encoding/json"
	"net/url"
	"fmt"
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
	query := url.QueryEscape(trend)
	n := 100
	responseBytes := makeAuthedRequest("GET", "1.1/search/tweets.json?q="+query+"&include_entities=false&lang=en&result_type=popular&count="+string(n))

	var parsed tweetsResponse
	json.Unmarshal(responseBytes, &parsed)
	tweets := extractTweetsfromResponse(parsed)

	nextTweets := getNextTweets(parsed.Search_metadata.Next_results, n-len(tweets), trend)
	tweets = append(tweets, nextTweets...)

	return tweets
}

//
// FetchTrends goes to twitter and returns a list of the top 5 trending topics
// the woeid of 2459115 hard codes this to the New York
//
func FetchTrends() []string {
	responseBytes := makeAuthedRequest("GET", "1.1/trends/place.json?id=2459115")

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

//
// Tweet post a tweet to the Twitter
///
func Tweet(tweet string) {
	tweet = url.QueryEscape(tweet)
	// responseBytes := makeAuthedRequest("POST", "1.1/statuses/update.json?status=" + tweet)

	fmt.Println(tweet)
}
