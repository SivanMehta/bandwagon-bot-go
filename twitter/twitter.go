package twitter

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strings"
)

// TBD: getting actual tweets from Twitter
var accessToken = ""
var token = ""
var secret = ""

const apiBase = "https://api.twitter.com/"

//
// GetTweets will return a body of text from trending topics
// This is currently just the text of the Declaration of Independence, but stay tuned.
//
func GetTweets(trend string) []string {
	// exactly containing the trend, without tweets that contain links
	query := "\"" + trend + "\" -filter:links"
	query = url.QueryEscape(query)
	// responseString := makeAuthedRequest("GET", "1.1/search/tweets.json?q="+query+"&include_entities=false&lang=en&result_type=popular")
	// fmt.Println(string(responseString))

	totallyTweets, _ := ioutil.ReadFile("corpus.txt")

	tweets := strings.Split(string(totallyTweets), ".")
	return tweets
}

//
// GetTrends goes to twitter and returns a list of trending topics
// the woeid of 2352824 hard codes this to the US
//
func GetTrends() []string {
	responseBytes := makeAuthedRequest("GET", "1.1/trends/place.json?id=2352824")

	var parsed trendResponse
	json.Unmarshal(responseBytes, &parsed)

	data := parsed[0].Trends
	trends := make([]string, len(data))

	for i := range data {
		trends[i] = data[i].Name
	}

	return trends
}
