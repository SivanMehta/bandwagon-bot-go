package chain

import (
	"math/rand"
	"io/ioutil"
)

var corpus []byte

//
// GenerateChain will return a random piece of the Declaration of the Independence
//
func GenerateChain() []byte {
	if corpus == nil {
		corpus, _ = ioutil.ReadFile("corpus.txt")
	}
	
  // we will randomly slice this later
  start := rand.Intn(len(corpus))
  end := start + rand.Intn(len(corpus) - start)
	return corpus[start:end]
}