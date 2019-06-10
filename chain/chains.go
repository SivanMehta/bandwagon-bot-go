package chain

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

type bandwagon func() string

// Key: Trending Topic on Twitter
// Value: Function that generates a tweet from the keyed topic
type bandwagons map[string]bandwagon

var currentBandwagons bandwagons
var upcomingBandwagons bandwagons

// character limit for tweets
const limit = 240
const debug = "-----debug------"

func selectFrom(choices []string) string {
	if len(choices) == 0 {
		return debug
	}
	return choices[rand.Intn(len(choices))]
}

func selectStarter(starters []bigram) bigram {
	if len(starters) == 0 {
		return bigram{debug, debug}
	}
	return starters[rand.Intn(len(starters))]
}

func createGenerator(trend string, pointer *bandwagon, pool *sync.WaitGroup) {
	fmt.Println("Generating Markov chain for", trend)

	tweets := getTweets(trend)
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

		return composedTweet + "\n"
	}

	pool.Done() // decrement the counter
}

//
// MakeChains will be run on an interval to generate chain from a given set of tweets
//
func MakeChains() {
	// allocate a map for the new trends
	upcomingBandwagons = make(bandwagons)
	trends := getTrends()

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
