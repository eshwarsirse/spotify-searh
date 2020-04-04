package app

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"spotify-search/provider"
	"spotify-search/services"
	"strconv"
)

type Application struct {
	router    *mux.Router
	searchSVC provider.SearchAPI
}

func CreateApplication() Application {
	a := Application{}
	a.load()
	return a
}

func (a *Application) load() {
	a.loadRoutes()

	searchService, err := services.NewSearchAPI()
	if err != nil {
		panic(err)
	}
	a.searchSVC = searchService
}

func (a *Application) loadRoutes() {
	a.router = mux.NewRouter().StrictSlash(true)
	a.addRoutes()
}

func (a *Application) addRoutes() {
	a.addIndex()
	a.addHealthChecks()
	a.addSearchRoute()
}

func (a *Application) StartServer(port int) error {
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(a.router)

	p := strconv.Itoa(port)
	log.Println("listing on port ", p)
	return http.ListenAndServe(":"+p, n)
}
