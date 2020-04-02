package app

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"strconv"
)

type Application struct {
	router        *mux.Router
}

func CreateApplication() Application {
	a := Application{}
	a.load()
	return a
}

func (a *Application) load() {
	a.loadRoutes()
}

func (a *Application) loadRoutes() {
	a.router = mux.NewRouter().StrictSlash(true)
	a.addRoutes()
}

func (a *Application) addRoutes() {
	a.addHealthChecks()
}

func (a *Application) StartServer(port int) error {
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(a.router)

	p := strconv.Itoa(port)
	log.Println("listing on port ", p)
	return http.ListenAndServe(":"+p, n)
}
