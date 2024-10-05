package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer{
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router:= mux.NewRouter()
	subRouter:= router.PathPrefix("/api/v1").Subrouter()

	// prevent not used
	_ = subRouter

	log.Println("Listening on", s.addr)


	return http.ListenAndServe(s.addr, router)
}