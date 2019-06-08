package chain

import (
	"math/rand"
	"io/ioutil"
)

var corpus []byte

// GenerateChainFrom return a random subset of the corpus
func GenerateChain() []byte {
	if corpus == nil {
		corpus, _ = ioutil.ReadFile("corpus.txt")
	}
	
  // we will randomly slice this later
  start := rand.Intn(len(corpus))
  end := start + rand.Intn(len(corpus) - start)
	return corpus[start:end]
}