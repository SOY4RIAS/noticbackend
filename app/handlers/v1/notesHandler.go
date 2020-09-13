package v1

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"noticbackend/app/models"
	"noticbackend/utils/response"

	"noticbackend/app/services/note"
	"noticbackend/config"
)

type NotesHandler struct {
	service note.Service
	config  *config.Config
}

func NotesRouter(s note.Service, c *config.Config, r *chi.Mux) {
	handler := &NotesHandler{service: s, config: c}

	r.Route(BaseRoute+"/notes", func(r chi.Router) {
		r.Get("/", handler.findAll)
		r.Post("/", handler.createNote)
	})
}

func (handler *NotesHandler) createNote(w http.ResponseWriter, r *http.Request) {
	request := new(models.Note)

	defer func() {
		_ = r.Body.Close()
	}()

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&request); err != nil {
		response.AsErrorJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	request.Initialize()

	if err := handler.service.Create(r.Context(), request); err != nil {
		response.AsErrorJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	payload := map[string]string{
		"message": "Successful created",
	}

	response.AsJson(w, http.StatusOK, payload)
}

func (handler *NotesHandler) findAll(w http.ResponseWriter, r *http.Request) {

	notes, err := handler.service.FindAll(r.Context())

	if err != nil {
		response.AsErrorJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.AsJson(w, http.StatusOK, notes)
}
