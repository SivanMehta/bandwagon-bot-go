package chain

import (
	"io/ioutil"
	"strings"
)

// TBD: getting actual tweets from Twitter
var accessToken = ""

const token = ""
const secret = ""

// will be removed when moving to actual tweets, for now
// this is just a cached version of corpus.txt
var totallyTweets []byte

//
// getTweets will return a body of text from trending topics
// This is currently just the text of the Declaration of Independence, but stay tuned.
//
func getTweets(trend string) []string {
	if totallyTweets == nil {
		// this in place of getting from twitter, instead
		// just reading from a text file
		totallyTweets, _ = ioutil.ReadFile("corpus.txt")
	}

	tweets := strings.Split(string(totallyTweets), ".")
	return tweets
}

func getTrends() []string {
	// would normally be fetched from twitter, just hardcoded for now
	return []string{"anything", "at", "all"}
}
