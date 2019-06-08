package main

import (
  "fmt"
  "net/http"
  "os"
  "math/rand"
  "io/ioutil"
)

func generateChainFrom (text []byte) []byte {
  // we will randomly slice this later

  start := rand.Intn(len(text))
  end := start + rand.Intn(len(text) - start)
	return text[start:end]
}

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "5000"
  }

  // would normally be pulled form twitter
  corpus, _ := ioutil.ReadFile("corpus.txt")

  const indexPage = "public/index.html"
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, indexPage)
  })

  http.HandleFunc("/sentence", func(w http.ResponseWriter, r *http.Request) {
    w.Write(generateChainFrom(corpus))
  })

  fmt.Printf("Listening on port %s\n\n", port)
  http.ListenAndServe(":" + port, nil)
}
