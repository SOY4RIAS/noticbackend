package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"net/http"
	v1 "noticbackend/app/handlers/v1"
	"noticbackend/app/services/note"
	"noticbackend/config"
	"noticbackend/database"
)

func main() {
	c := config.New()

	db := database.New(c)

	notesService := note.ServiceNote{}.New(db, c)

	r := chi.NewRouter()

	setupMiddlewares(r)

	v1.NotesRouter(notesService, c, r)

	_ = http.ListenAndServe(c.Address, r)

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
