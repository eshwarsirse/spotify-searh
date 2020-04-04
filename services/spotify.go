package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"spotify-search/models"
	"spotify-search/provider"
	"strconv"
)

const (
	SpotifyBaseURL    = "https://api.spotify.com"
	SpotifySearchPath = "/v1/search"
)

type searchAPI struct {
	token         *models.Token
	authenticator provider.Authenticator
}

func NewSearchAPI() (provider.SearchAPI, error) {
	authenticator := NewAuthenticator()
	token, err := authenticator.GenerateAccessToken()

	if err != nil {
		return nil, err
	}

	return &searchAPI{token: token, authenticator: authenticator}, nil
}

func (s *searchAPI) Search(query string, searchtype string, limit int, offset int) ([]byte, error) {

	searchURL, err := url.Parse(SpotifyBaseURL)
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
		token, err := s.authenticator.GenerateAccessToken()
		if err != nil {
			return err
		}
		s.token = token
	}
	return nil
}
