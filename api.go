package main

import (
	"log"
	"net/http"
	"time"
)

type APIServer struct {
	addr        string
	Controllers []Controller
}

func NewAPIServer(addr string, controllers ...Controller) *APIServer {
	return &APIServer{addr: addr, Controllers: controllers}
}

func (s *APIServer) Start() error {
	mux := http.NewServeMux()
	for _, controller := range s.Controllers {
		controller.RegisterRoutes(mux)
	}
	server := http.Server{
		Addr:         s.addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("Starting server on", s.addr)
	return server.ListenAndServe()
}
