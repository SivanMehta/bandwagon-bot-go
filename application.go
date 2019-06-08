package main

import (
  "fmt"
  "net/http"
  "os"

  "./chain"
)

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "5000"
  }

  const indexPage = "public/index.html"
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, indexPage)
  })

  http.HandleFunc("/sentence", func(w http.ResponseWriter, r *http.Request) {
    w.Write(chain.GenerateChain())
  })

  fmt.Printf("Listening on port %s\n\n", port)
  http.ListenAndServe(":" + port, nil)
}
