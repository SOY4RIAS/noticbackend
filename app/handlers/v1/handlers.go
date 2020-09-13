package v1

import (
	"github.com/go-chi/chi"
	note2 "noticbackend/app/handlers/v1/note"

	"noticbackend/app/services/note"
	"noticbackend/config"
	"noticbackend/database"
)

const (
	BaseRoute = "/api/v1"
)

func Setup(c *config.Config, r *chi.Mux) {

	db := database.New(c)

	serviceNote := note.ServiceNote{}.New(db, c)

	note2.NotesRouter(BaseRoute, serviceNote, c, r)
}
