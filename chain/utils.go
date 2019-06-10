package chain

import (
	"math/rand"
	"strings"
)

// we're hard-coding order of 2 markov chains here
type bigram struct {
	First, Second string
}
type chain map[bigram][]string

type bandwagon func() string
type bandwagons map[string]bandwagon

// character limit for tweets
const limit = 240

// this is a uuid that is very unlikely to ever be tweeted
const debug = "d99c23df-cc98-4660-9b8c-0dede4741c79"

func selectFrom(choices []string) string {
	if len(choices) == 0 {
		return debug
	}
	return choices[rand.Intn(len(choices))]
}

// the same as above. but for bigrams (beceause go hates generics)
func selectStarter(starters []bigram) bigram {
	if len(starters) == 0 {
		return bigram{debug, debug}
	}
	return starters[rand.Intn(len(starters))]
}

// Use the pair of words at the beginning of each tweet as a set of
// potentional starters
func findStarters(tweets []string) []bigram {
	starters := []bigram{}

	for _, tweet := range tweets {
		words := strings.Split(tweet, " ")

		if len(words) > 2 {
			pair := bigram{words[0], words[1]}
			starters = append(starters, pair)
		}
	}

	return starters
}

// concatenate some strings
func join(strs ...string) string {
	var ret string
	for _, str := range strs {
		ret += str
	}
	return ret
}
