package twitter

import (
	"io/ioutil"
	"os"
	"strings"
)

// TBD: getting actual tweets from Twitter
var accessToken = ""
var token = ""
var secret = ""

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
	// would normally be fetched from twitter, just hardcoded for now
	return []string{"anything", "at", "all"}
}

//
// GenerateAccessToken reads the secrets from the environment variables and
// uses them to fetch an access token
func GenerateAccessToken() {
	token = os.Getenv("TOKEN")
	secret = os.Getenv("SECRET")
	accessToken = "12345"
}
