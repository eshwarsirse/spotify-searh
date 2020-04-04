package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

//Search accepts query, type, limit(optional), offset(optional) for each request.
func (a *Application) Search(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	searchQuery := queryValues.Get("query")
	if len(searchQuery) <= 0 {
		a.writeError(w, "Missing 'query' param")
		return
	}

	searchType := queryValues.Get("type")
	if len(searchType) <= 0 {
		a.writeError(w, "Missing 'type' param")
		return
	}

	var searchLimit, searchOffset int
	var err error

	limit := queryValues.Get("limit")
	if len(limit) > 0 {
		searchLimit, err = strconv.Atoi(limit)
		if err != nil {
			a.writeError(w, "Invalid 'limit' param")
			return
		}
	}

	offset := queryValues.Get("offset")
	if len(offset) > 0 {
		searchOffset, err = strconv.Atoi(offset)
		if err != nil {
			a.writeError(w, "Invalid 'offset' param")
			return
		}
	}

	resp, err := a.searchSVC.Search(searchQuery, searchType, searchLimit, searchOffset)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (a *Application) writeError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	err := struct {
		Description string `json:"description"`
	}{
		msg,
	}
	_ = json.NewEncoder(w).Encode(err)
}

func (a *Application) addSearchRoute() {
	a.router.HandleFunc("/search", a.Search)
}
