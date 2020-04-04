package services

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"spotify-search/models"
	"spotify-search/provider"
	"strings"
	"time"
)

const (
	SpotifyLoginURL  = "https://accounts.spotify.com"
	SpotifyLoginPath = "/api/token"
)

type authenticator struct {
}

func NewAuthenticator() provider.Authenticator {
	return &authenticator{}
}

func (a *authenticator) GenerateAccessToken() (*models.Token, error) {
	loginURL, err := url.Parse(SpotifyLoginURL)
	if err != nil {
		return nil, err
	}

	loginURL.Path = SpotifyLoginPath
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	basicToken := base64.StdEncoding.EncodeToString([]byte(provider.ClientID + ":" + provider.ClientSecret))

	req, err := http.NewRequest("POST", loginURL.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", basicToken))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	token := &models.Token{}
	err = json.NewDecoder(resp.Body).Decode(token)
	if err != nil {
		return token, err
	}
	log.Println("New Access token generated successfully")
	token.Expiry = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	return token, err
}
