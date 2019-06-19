package twitter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

const apiBase = "https://api.twitter.com/"

// make an authed request to the twitter API
func makeAuthedRequest(method string, url string) []byte {
	body := []byte("")

	req, _ := http.NewRequest(method, apiBase+url, bytes.NewBuffer(body))
	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic("Could not make call to " + url)
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	return bodyBytes
}

//
// GenerateAccessToken reads the secrets from the environment variables and
// uses them to fetch an access token
func GenerateAccessToken() {
	token = os.Getenv("KEY")
	secret = os.Getenv("SECRET")

	if len(token) == 0 || len(secret) == 0 {
		message := `
		Could not find environment variables for the twitter API, make sure they're
		available per the README
		`

		panic(message)
	}

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
}
