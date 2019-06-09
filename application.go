package main

import (
  "fmt"
  "net/http"
  "os"
  "time"
  "strings"

  "./chain"
)

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "5000"
  }

  // make a chain before setting the routes
  // so that they're available on server start
  chain.MakeChains()

  const indexPage = "public/index.html"
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, indexPage)
  })

  http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
    trend := strings.Replace(r.URL.String(), "/api/", "", -1)
    tweet := chain.TweetFromTrend(trend)
    w.Write([]byte(tweet))
  })

  // Regenerate the chains on an interval
  // to update what bandwagon we're hopping on.
  go func() {
    for true {
      time.Sleep(time.Second * 5)
      chain.MakeChains()
    }
  }()

  fmt.Printf("Listening on port %s\n", port)
  http.ListenAndServe(":" + port, nil)
}
