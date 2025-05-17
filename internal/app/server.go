package server

import (
	"log"
	"net/http"
)

const (
	host            = "localhost"
	port            = "8080"
	letters         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	shortenerLength = 8
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s Server) Run(addr string) error {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", shortURL)
	mux.HandleFunc("/{id}", getURL)

	log.Printf("Starting server at %s...\r\n", addr)
	return http.ListenAndServe(addr, mux)
}
