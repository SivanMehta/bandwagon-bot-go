package chain

import (
	"math/rand"
)

//
// GenerateChain will return a random piece of the gathered tweets
//
func GenerateChain() []byte {
	corpus := getTweets()
  // we will randomly slice this later
  start := rand.Intn(len(corpus))
  end := start + rand.Intn(len(corpus) - start)
	return corpus[start:end]
}