package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"net/http"
	v1 "noticbackend/app/handlers/v1"
	"noticbackend/config"
)

type server struct {
	config *config.Config
	router *chi.Mux
}

func New() *server {
	c := config.New()
	r := chi.NewRouter()

	return &server{config: c, router: r}
}

func (s *server) ListenAndServe() error {
	setupMiddlewares(s.router)

	v1.Setup(s.config, s.router)

	return http.ListenAndServe(s.config.Address, s.router)
}

func setupMiddlewares(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	setupCors(r)
}

func setupCors(r *chi.Mux) {
	options := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}

	r.Use(cors.Handler(options))
}
