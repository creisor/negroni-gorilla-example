package main

import (
	"github.com/codegangsta/negroni"
	"github.com/creisor/negroni-gorilla-example/v1"
	"github.com/creisor/negroni-gorilla-example/v2"
	"github.com/gorilla/mux"
)

func main() {
	v1Router := mux.NewRouter()
	v1 := v1.NewServer(v1Router, "Hello from v1\n")

	v2Router := mux.NewRouter().PathPrefix("/v2/").Subrouter()
	v2 := v2.NewServer(v2Router, "Hello from v2\n")

	v1.Routes()
	v2.Routes()

	// Order matters!  The "/" route will match anything, so the versioned routes
	// should go first.
	router := mux.NewRouter()
	router.PathPrefix("/v2").Handler(v2Router)
	router.PathPrefix("/").Handler(v1Router)

	n := negroni.Classic()
	n.UseHandler(router)

	n.Run(":3000")
}
