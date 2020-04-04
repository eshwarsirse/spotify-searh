package services

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/url"
	"spotify-search/models"
	"spotify-search/provider"
	"strings"
	"time"
)

const (
	SpotifyLoginPath = "/api/token"
)

type authenticator struct {
	client *http.Client
	config *viper.Viper
}

//NewAuthenticator will return an implementation of Authenticator interface
func NewAuthenticator(config *viper.Viper) provider.Authenticator {
	return &authenticator{client: &http.Client{}, config: config}
}

//GenerateAccessToken gets access token for given app credentials using below flow
//https://developer.spotify.com/documentation/general/guides/authorization-guide/#client-credentials-flow
func (a *authenticator) GenerateAccessToken() (*models.Token, error) {
	loginURL, err := url.Parse(a.config.GetString("SPOTIFY_ACCOUNTS_BASE_URL"))
	if err != nil {
		return nil, err
	}

	loginURL.Path = SpotifyLoginPath
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", loginURL.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	clientID := a.config.GetString("SPOTIFY_CLIENT_ID")
	clientSecret := a.config.GetString("SPOTIFY_CLIENT_SECRET")
	basicToken := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", basicToken))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := a.client.Do(req)
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
