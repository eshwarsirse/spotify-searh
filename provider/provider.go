package provider

import "spotify-search/models"

const (
	ClientID     = "14f3d2a35fd74a3983f42840fcf18580"
	ClientSecret = "e3ec579edf404643966e400fd211505a"
)

//SearchAPI is used to search artist,albums, tracks based on search query
type SearchAPI interface {
	Search(query string, searchType string, limit int, offset int) ([]byte, error)
}

//Authenticator is used to authenticate app using client credentials
type Authenticator interface {
	GenerateAccessToken() (*models.Token, error)
}
