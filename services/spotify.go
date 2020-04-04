package services

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"spotify-search/models"
	"spotify-search/provider"
	"strconv"
)

const (
	SpotifySearchPath = "/v1/search"
)

type searchAPI struct {
	token         *models.Token
	authenticator provider.Authenticator
	config *viper.Viper
}

//NewSearchAPI returns an implementation of SearchAPI interface
func NewSearchAPI(config *viper.Viper) (provider.SearchAPI, error) {
	authenticator := NewAuthenticator(config)
	token, err := authenticator.GenerateAccessToken()

	if err != nil {
		return nil, err
	}

	return &searchAPI{token: token, authenticator: authenticator, config: config}, nil
}

//Search implements Spotify search API
//https://developer.spotify.com/documentation/web-api/reference/search/search/
func (s *searchAPI) Search(query string, searchtype string, limit int, offset int) ([]byte, error) {

	searchURL, err := url.Parse(s.config.GetString("SPOTIFY_BASE_URL"))
	if err != nil {
		return nil, err
	}

	searchURL.Path = SpotifySearchPath
	params := url.Values{}
	params.Set("q", query)
	params.Set("type", searchtype)
	if limit > 0 {
		params.Set("limit", strconv.Itoa(limit))
	}
	if offset > 0 {
		params.Set("offset", strconv.Itoa(offset))
	}
	searchURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", searchURL.String(), nil)
	if err != nil {
		return nil, err
	}

	err = s.validateToken()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token.AccessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(data))
	}
	return data, err
}

func (s *searchAPI) validateToken() error {
	if s.token.IsExpired() {
		log.Println("AccessToken expired, renewing it..")
		token, err := s.authenticator.GenerateAccessToken()
		if err != nil {
			return err
		}
		s.token = token
	}
	return nil
}
