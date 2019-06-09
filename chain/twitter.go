package chain

import (
  "io/ioutil"
  "fmt"
)

var corpus []byte

// TBD: getting an actual corpus from Twitter
var accessToken = ""
const token = ""
const secret = ""

//
// getTweets will return a body of text from trending topics
// This is currently just the text of the Declaration of Independence, but stay tuned.
// 
func getTweets() []byte {
  if corpus == nil {
    fmt.Println("reading corpus")
    // this in place of getting from twitter, instead
    // just reading from a text file
    corpus, _ = ioutil.ReadFile("corpus.txt")
  }

  return corpus
}

func getTrends() []string {
  // would normally be fetched from twitter, just hardcoded for now
  return []string{"anything", "at", "all"}
}
