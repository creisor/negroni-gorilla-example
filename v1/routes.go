package v1

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	Router  *mux.Router
	message string
}

func (s *server) Routes() {
	s.Router.HandleFunc("/hello", s.handleV1()).Methods("HEAD", "GET")
}

func (s *server) handleV1() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, s.message)
		return
	}
}

func NewServer(r *mux.Router, message string) *server {
	s := &server{
		Router:  r,
		message: message,
	}
	return s
}
