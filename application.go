package main

import (
	"flag"
	"log"
	"spotify-search/app"
)

func main() {

	a := app.CreateApplication()
	var port = flag.Int("port", 8080, "port on which the server should run")

	err := a.StartServer(*port)
	if err != nil {
		log.Println("unable to start server")
	}
}
