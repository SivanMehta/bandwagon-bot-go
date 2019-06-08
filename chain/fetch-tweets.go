package chain

import (
	"io/ioutil"
)

var corpus []byte

//
// GetTweets will return a body of text from trending topics
// This is currently just the text of the Declaration of Independence, but stay tuned.
// 
func getTweets() []byte {
	if corpus == nil {
		corpus, _ = ioutil.ReadFile("corpus.txt")
	}

	return corpus
}