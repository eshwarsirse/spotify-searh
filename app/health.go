package app

import (
	"net/http"
)

//Liveness checks if service is reachable
func (a *Application) Liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

//Readiness checks if service can process requests
func (a *Application) Readiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (a *Application) addHealthChecks() {
	a.router.HandleFunc("/liveness", a.Liveness)
	a.router.HandleFunc("/readiness", a.Readiness)
}
