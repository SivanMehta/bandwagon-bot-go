package chain

import (
  "fmt"
  // "time"
  // "reflect"
)

// Key: Trending Topic on Twitter
// Value: Function that generates a tweet from the keyed topic
type bandwagons map[string]func() string
var currentBandwagons bandwagons
var upcomingBandwagons bandwagons

func makeChain(trend string) {
  fmt.Println("Generating Markov chain for", trend)

  // BUILD GATHER TWEETS FROM GIVEN TREND
  // BUILD A NGRAM DICTIONARY FROM THE TWEETS
  // ASSIGN A FUNCTION THAT USES THAT DICTIONARY
  upcomingBandwagons[trend] = func() string {
    return "Raptors 4 life"
  }
}

//
// MakeChains will be run on an interval to generate chain from a given set of tweets
//
func MakeChains() {
  // allocate a map for the new trends
  upcomingBandwagons = make(bandwagons)
  trends := getTrends()

  // figure how to make concurrent, because that will become blocking if done poorly
  for _, trend := range trends {
    makeChain(trend)
  }

  
  // reset what bandwagon we're currently on
  currentBandwagons = upcomingBandwagons
}

// concatenate some strings
func join(strs ...string) string {
	var ret string
	for _, str := range strs {
		ret += str
	}
  return ret
  
}

//
// TweetFromTrend will try and generate a tweet from the given trend
//
func TweetFromTrend(trend string) string {
  if generator, ok := currentBandwagons[trend]; ok {
    return generator()
  }
  return join(trend, " is a not a currently available bandwagon")
}