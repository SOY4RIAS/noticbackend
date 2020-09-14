package note

import (
	"github.com/go-chi/chi"
	"noticbackend/app/services/note"
	"noticbackend/config"
)

func NotesRouter(prefix string, s note.Service, c *config.Config, r *chi.Mux) {
	handler := &NotesHandler{service: s, config: c}

	r.Route(prefix+"/notes", func(r chi.Router) {
		r.Get("/", handler.findAll)
		r.Get("/{noteID}", handler.findOneById)
		r.Put("/{noteID}", handler.update)
		r.Delete("/{noteID}", handler.delete)
		r.Post("/", handler.createNote)
	})
}
