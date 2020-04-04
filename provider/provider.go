package provider

import "spotify-search/models"

const (
	ClientID     = "14f3d2a35fd74a3983f42840fcf18580"
	ClientSecret = "e3ec579edf404643966e400fd211505a"
)

type SearchAPI interface {
	Search(query string, searchtype string, limit int, offset int) ([]byte, error)
}

type Authenticator interface {
	GenerateAccessToken() (*models.Token, error)
}
