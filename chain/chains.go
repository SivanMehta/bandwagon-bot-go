package chain

import (
	"fmt"
	"strings"
	"sync"

	"../twitter"
)

func createGenerator(trend string, pointer *bandwagon, pool *sync.WaitGroup) {
	fmt.Println("Generating Markov chain for", trend)

	tweets := twitter.GetTweets(trend)
	dictionary := buildDictionary(tweets)
	startingPairs := findStarters(tweets)

	// Gather tweets from a given trend
	// Build an n-gram dictionary from the tweets
	// Generate a function that uses that dictionary to return a tweet
	*pointer = func() string {
		startingPair := selectStarter(startingPairs)
		tweet := []string{startingPair.First + " ", startingPair.Second}
		length := len(tweet[0]) + len(tweet[1])
		currentKey := startingPair

		for true {
			if candidates, ok := dictionary[currentKey]; ok {
				choice := selectFrom(candidates)

				if len(choice)+length > limit {
					// This composition is longer than twitter will allow
					break
				}

				tweet = append(tweet, " "+choice)
				length += len(choice)
				currentKey = bigram{currentKey.Second, choice}
			} else {
				// we have reached a n-gram that has never been followed before, so
				// we just follow past advice and just stop here
				break
			}
		}

		composedTweet := join(tweet...)
		composedTweet = strings.Replace(composedTweet, debug, "", -1)

		return composedTweet
	}

	pool.Done() // decrement the counter
}

var currentBandwagons bandwagons
var upcomingBandwagons bandwagons

//
// MakeChains will be run on an interval to generate chain from a given set of tweets
//
func MakeChains() {
	// allocate a map for the new trends
	upcomingBandwagons = make(bandwagons)
	trends := twitter.GetTrends()

	// a pool so that we can generate the bandwagons in parallel
	var pool sync.WaitGroup
	// an array of pointers that we'll use to track the generated trends
	generators := make([]bandwagon, len(trends))
	for i, trend := range trends {
		pool.Add(1) // increment the counter
		go createGenerator(trend, &generators[i], &pool)
	}
	pool.Wait() // wait until the counter is at 0
	// now we have two analogous arrays, trends and generators that represent the upcoming bandwagon
	// so we can very quickly arrange build the upcomingBandwagon map
	for i, trend := range trends {
		upcomingBandwagons[trend] = generators[i]
	}

	// reset what bandwagon we're currently on
	currentBandwagons = upcomingBandwagons
}

//
// FromTrend will try and generate a tweet from the given trend
//
func FromTrend(trend string) string {
	if generator, ok := currentBandwagons[trend]; ok {
		return join(generator(), "\n")
	}
	return join("\"", trend, "\" is a not a currently available bandwagon\n")
}
