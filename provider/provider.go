package provider

import "spotify-search/models"

//SearchAPI is used to search artist,albums, tracks based on search query
type SearchAPI interface {
	Search(query string, searchType string, limit int, offset int) ([]byte, error)
}

//Authenticator is used to authenticate app using client credentials
type Authenticator interface {
	GenerateAccessToken() (*models.Token, error)
}
