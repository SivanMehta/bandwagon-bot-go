package twitter

import (
	"fmt"
	"io/ioutil"
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
	totallyTweets, _ := ioutil.ReadFile("corpus.txt")

	tweets := strings.Split(string(totallyTweets), ".")
	return tweets
}

//
// GetTrends goes to twitter and returns a list of trending topics in the US
//
func GetTrends() []string {
	responseString := makeAuthedRequest("GET", "1.1/trends/place.json?id=1")

	fmt.Println(responseString)

	// would normally be fetched from twitter, just hardcoded for now
	return []string{"anything", "at", "all"}
}
