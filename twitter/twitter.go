package twitter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// TBD: getting actual tweets from Twitter
var accessToken = ""
var token = ""
var secret = ""

const apiBase = "https://api.twitter.com/"

//
// GetTweets will return a body of text from trending topics
// This is currently just the text of the Declaration of Independence, but stay tuned.
//
func GetTweets(trend string) []string {
	totallyTweets, _ := ioutil.ReadFile("corpus.txt")

	tweets := strings.Split(string(totallyTweets), ".")
	return tweets
}

//
// GetTrends goes to twitter and returns a list of trending topics in the US
//
func GetTrends() []string {
	// would normally be fetched from twitter, just hardcoded for now
	return []string{"anything", "at", "all"}
}

//
// GenerateAccessToken reads the secrets from the environment variables and
// uses them to fetch an access token
func GenerateAccessToken() {
	token = os.Getenv("KEY")
	secret = os.Getenv("SECRET")

	body := []byte("grant_type=client_credentials")

	req, err1 := http.NewRequest("POST", apiBase+"oauth2/token", bytes.NewBuffer(body))
	if err1 != nil {
		panic("Could not fetch access token")
	}

	msg := token + ":" + secret
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Authorization", "Basic "+encoded)

	client := http.Client{}
	res, err2 := client.Do(req)
	if err2 != nil {
		panic("Could not fetch access token")
	}

	var target map[string]string
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&target)

	accessToken = target["access_token"]

	fmt.Println(accessToken)
}
