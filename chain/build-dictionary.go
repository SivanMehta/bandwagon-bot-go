package chain

import (
	"sync"
	"strings"
)

// we're hard-coding order of 1 markov chains here
type chain map[string][]string

func chainFromTweet(tweet string, pointer *chain, pool *sync.WaitGroup) {
	tokens := strings.Split(tweet, " ")
	tweetChain := make(chain)

	for i := 1; i < len(tokens); i ++ {
		if arr, ok := tweetChain[tokens[i - 1]]; ok {
			tweetChain[tokens[i - 1]] = append(arr, tokens[i])
		} else {
			tweetChain[tokens[i - 1]] = []string{tokens[i]}
		}
	}

	*pointer = tweetChain
	pool.Done()
}

// given a list of tweets, generate a dictionary
func buildDictionary(tweets []string) chain {
	var pool sync.WaitGroup
	chains := make([]chain, len(tweets))

	for i, tweet := range tweets {
		pool.Add(1)
		go chainFromTweet(tweet, &chains[i], &pool)
	}
	pool.Wait()

	// now combine the dictionaries for each tweet into one giant dictionary for
	// the corpus of tweets. We build one individually for each tweet and then
	// combine instead of doing it all at once because we don't want to
	// erronously generate a chain from tweets that happen to be adjacent

	return chains[0]
}

// Use the first word of each tweet as a baseline starter
func findStarters(tweets []string) []string {
	starters := []string{}

	for _, tweet := range(tweets) {
		starters = append(starters, strings.Split(tweet, " ")[0])
	}

	return starters
}
