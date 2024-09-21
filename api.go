package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gotaskmaster/routes"
	"log"
	"net/http"
)

type Config struct {
	Addr string
}

type Server struct {
	r      *chi.Mux
	config Config
}

func NewServer(cfg Config) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	apiRouter := chi.NewRouter()
	apiRouter.Mount("/auth", routes.HandleAuth())

	r.Mount("/api", apiRouter)
	return &Server{
		r:      r,
		config: cfg,
	}
}

func (s *Server) Start() error {
	if s.config.Addr == "" {
		s.config.Addr = ":8080"
	}
	log.Printf("Listening on %s", s.config.Addr)

	err := http.ListenAndServe(s.config.Addr, s.r)
	return err
}
