package server

import (
	"log"
	"net/http"
	"os"
)

type Server struct {
	port   string
	webDir string
}

func New() *Server {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	return &Server{
		port:   port,
		webDir: "./web",
	}
}

func (s *Server) Start() error {
	fs := http.FileServer(http.Dir(s.webDir))
	http.Handle("/", fs)

	log.Printf("Starting web server on port %s", s.port)
	return http.ListenAndServe(":"+s.port, nil)
}
