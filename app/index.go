package app

import (
	"encoding/json"
	"net/http"
)

//Index shows default page/welcome page
func (a *Application) Index(w http.ResponseWriter, r *http.Request) {

	resp := struct {
		Msg string `json:"msg"`
	}{
		"Welcome to Spotify Search Service",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (a *Application) addIndex() {
	a.router.HandleFunc("/", a.Index)
}
