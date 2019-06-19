package chain

import (
	"strings"
	"sync"
)

func chainFromTweet(tweet string, pointer *chain, pool *sync.WaitGroup) {
	tokens := strings.Split(tweet, " ")
	tweetChain := make(chain)

	for i := 2; i < len(tokens); i++ {
		key := bigram{tokens[i-2], tokens[i-1]}
		if arr, ok := tweetChain[key]; ok {
			tweetChain[key] = append(arr, tokens[i])
		} else {
			tweetChain[key] = []string{tokens[i]}
		}
	}

	*pointer = tweetChain
	pool.Done()
}

// given a list of tweets, generate a dictionary
func buildDictionary(tweets []string) chain {
	var pool sync.WaitGroup
	tweetChains := make([]chain, len(tweets))

	for i, tweet := range tweets {
		pool.Add(1)
		go chainFromTweet(tweet, &tweetChains[i], &pool)
	}
	pool.Wait()

	// now combine the dictionaries for each tweet into one giant dictionary for
	// the corpus of tweets. We build one individually for each tweet and then
	// combine instead of doing it all at once because we don't want to
	// erronously generate a chain from tweets that happen to be adjacent

	markovChain := make(chain)
	for _, tweetChain := range tweetChains {
		for key := range tweetChain {
			if arr, ok := markovChain[key]; ok {
				markovChain[key] = append(arr, tweetChain[key]...)
			} else {
				markovChain[key] = tweetChain[key]
			}
		}
	}

	return markovChain
}
