package app

import (
	"github.com/spf13/viper"
)

func (a *Application) loadConfig() {
	c := viper.New()
	c.SetDefault("SPOTIFY_BASE_URL", "https://api.spotify.com")
	c.SetDefault("SPOTIFY_ACCOUNTS_BASE_URL", "https://accounts.spotify.com")
	c.SetDefault("SPOTIFY_CLIENT_ID", "14f3d2a35fd74a3983f42840fcf18580")
	c.SetDefault("SPOTIFY_CLIENT_SECRET", "e3ec579edf404643966e400fd211505a")
	a.config = c
	c.AutomaticEnv()
}

