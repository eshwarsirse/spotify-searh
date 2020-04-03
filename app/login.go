package app

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	ClientID = "14f3d2a35fd74a3983f42840fcf18580"
	ClientSecret = "e3ec579edf404643966e400fd211505a"
)

func (a *Application) Login(w http.ResponseWriter, r *http.Request) {

	authURL, err := url.Parse("https://accounts.spotify.com")
	if err != nil {
		log.Println("Malformed URL")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	authURL.Path = "authorize"

	values := url.Values{}
	values.Add("client_id", ClientID)
	values.Add("response_type", "code")
	values.Add("scope", "")
	values.Add("redirect_uri", "")
	values.Add("state", "")

	authURL.RawQuery = values.Encode()

	log.Println("Encoded URL: ", authURL.String())
	http.Redirect(w, r, authURL.String(), http.StatusPermanentRedirect)

}

func (a *Application) LoginV2(w http.ResponseWriter, r *http.Request) {
	loginURL, err := url.Parse("https://accounts.spotify.com")
	if err != nil {
		log.Println("Malformed URL")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	loginURL.Path = "/api/token"

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	log.Println("loginURL: ", loginURL.String())

	basicToken := base64.StdEncoding.EncodeToString([]byte(ClientID + ":" + ClientSecret))

	req, err := http.NewRequest("POST", loginURL.String(), strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", basicToken))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(resp.StatusCode)
	tokenData, _ := ioutil.ReadAll(resp.Body)
	log.Println("----------------------")
	log.Println(string(tokenData))
	log.Println("----------------------")
}

func (a *Application) addLoginRoutes() {
	a.router.HandleFunc("/login", a.LoginV2)
}