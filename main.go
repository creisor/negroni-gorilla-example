package main

import (
	"log"
	"net/http"

	"github.com/creisor/negroni-gorilla-example/v1"
	"github.com/creisor/negroni-gorilla-example/v2"
	"github.com/gorilla/mux"
)

func main() {
	v1Router := mux.NewRouter().PathPrefix("/v1").Subrouter()
	v1 := v1.NewServer(v1Router, "Hello from v1\n")

	v2Router := mux.NewRouter().PathPrefix("/v2").Subrouter()
	v2 := v2.NewServer(v2Router, "Hello from v2\n")

	v1.Routes()
	v2.Routes()

	http.Handle("/v1/", v1Router)
	http.Handle("/v2/", v2Router)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
