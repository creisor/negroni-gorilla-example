package main

import (
	"scratch/multiple_routes/v1"
	"scratch/multiple_routes/v2"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	v1Router := mux.NewRouter().PathPrefix("/v1").Subrouter()
	v1 := v1.NewServer(v1Router, "Hello from v1\n")

	v2Router := mux.NewRouter().PathPrefix("/v2").Subrouter()
	v2 := v2.NewServer(v2Router, "Hello from v2\n")

	v1.Routes()
	v2.Routes()

	n := negroni.Classic()
	n.UseHandler(v1.Router)
	n.UseHandler(v2.Router)

	n.Run(":3000")
}