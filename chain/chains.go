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

func selectFrom(choices []string) string {
	if len(choices) == 0 {
		return "-----debug------"
	}
	return choices[rand.Intn(len(choices))]
}

func createGenerator(trend string, pointer *bandwagon, pool *sync.WaitGroup) {
	fmt.Println("Generating Markov chain for", trend)

	tweets := getTweets(trend)
	dictionary := buildDictionary(tweets)
	starters := findStarters(tweets)

	// Gather tweets from a given trend
	// Build an n-gram dictionary from the tweets
	// Generate a function that uses that dictionary to return a tweet
	// (THIS IS THE HARD PART)
	*pointer = func() string {
		starter := selectFrom(starters)
		tweet := []string{starter}
		length := len(starter)
		current := starter

		for true {
			candidates := dictionary[current]
			choice := selectFrom(candidates)

			if len(choice)+length > 240 {
				break
			}

			tweet = append(tweet, " "+choice)
			length += len(choice)
			current = choice
		}

		composedTweet := join(tweet...)
		composedTweet = strings.Replace(composedTweet, "-----debug------", "", -1)

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
		return join(generator(), "\n")
	}
	return join("\"", trend, "\" is a not a currently available bandwagon\n")
}
