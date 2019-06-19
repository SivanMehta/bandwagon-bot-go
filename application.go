package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"./chain"
	"./twitter"
)

type HomePage struct {
	Trends []string
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// fetch the access token from the API so we can make authenticated requests
	twitter.GenerateAccessToken()

	// make a chain before setting the routes
	// so that they're available on server start
	chain.MakeChains()

	tmpl := template.Must(template.ParseFiles("public/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := HomePage{twitter.GetTrends()}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		trend := strings.Replace(r.URL.String(), "/api/", "", -1)
		parsedTrend, _ := url.PathUnescape(trend)
		tweet := chain.FromTrend(parsedTrend)
		w.Write([]byte(tweet))
	})

	// Regenerate the chains on an interval
	// to update what bandwagon we're hopping on.
	go func() {
		for true {
			time.Sleep(time.Hour)
			chain.MakeChains()
		}
	}()

	fmt.Printf("Listening on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
