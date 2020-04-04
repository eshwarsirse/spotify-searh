package app

import (
	"log"
	"net/http"
	"strconv"
)

func (a *Application) Search(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	searchQuery := queryValues.Get("query")
	if len(searchQuery) <= 0 {
		a.writeError(w, "missing query param")
		return
	}

	searchType := queryValues.Get("type")
	if len(searchType) <= 0 {
		a.writeError(w, "missing search type")
		return
	}

	var searchLimit, searchOffset int
	var err error

	limit := queryValues.Get("limit")
	if len(limit) > 0 {
		searchLimit, err = strconv.Atoi(limit)
		if err != nil {
			a.writeError(w, "invalid limit format")
			return
		}
	}

	offset := queryValues.Get("offset")
	if len(offset) > 0 {
		searchOffset, err = strconv.Atoi(offset)
		if err != nil {
			a.writeError(w, "invalid offset format")
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
	_, _ = w.Write([]byte(msg))
}

func (a *Application) addSearchRoute() {
	a.router.HandleFunc("/search", a.Search)
}
