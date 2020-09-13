package v1note2

import (
	"github.com/go-chi/chi"
	. "noticbackend/app/handlers/v1/note"

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

	NotesRouter(BaseRoute, serviceNote, c, r)
}
