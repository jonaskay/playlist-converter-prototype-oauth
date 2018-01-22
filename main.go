package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
	"google.golang.org/appengine"
)

type Credentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func main() {
	var c Credentials
	file, _ := ioutil.ReadFile("./credentials.json")
	json.Unmarshal(file, &c)

	spotifyOAuth2Config := &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		Endpoint:     spotify.Endpoint,
		RedirectURL:  "http://localhost:8080/auth/spotify/callback",
		Scopes:       []string{"playlist-modify-private"},
	}

	http.HandleFunc("/auth/spotify", func(w http.ResponseWriter, r *http.Request) {
		url := spotifyOAuth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})

	appengine.Main()
}
