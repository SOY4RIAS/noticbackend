package note

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"noticbackend/app/models"
	"noticbackend/app/services/note"
	"noticbackend/config"
	"noticbackend/utils/response"
)

type NotesHandler struct {
	service note.Service
	config  *config.Config
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

	payload := FindAll{
		Data: DataNotes{
			Notes: notes,
		},
	}

	response.AsJson(w, http.StatusOK, payload)
}

func (handler *NotesHandler) findOneById(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "noteID")

	noteFetched, err := handler.service.FindOneById(r.Context(), id)

	if err != nil {
		response.AsErrorJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	payload := FindOneById{
		Data: DataNote{
			Note: *noteFetched,
		},
	}

	response.AsJson(w, http.StatusOK, payload)
}

func (handler *NotesHandler) update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "noteID")
	noteUpdate := new(models.NoteUpdate)

	defer func() {
		_ = r.Body.Close()
	}()

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&noteUpdate); err != nil {
		response.AsErrorJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := handler.service.Update(r.Context(), id, *noteUpdate)

	if err != nil {
		response.AsErrorJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	payload := map[string]string{
		"message": "Successful updated",
	}

	response.AsJson(w, http.StatusOK, payload)
}

func (handler *NotesHandler) delete(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "noteID")

	err := handler.service.Delete(r.Context(), id)

	if err != nil {
		response.AsErrorJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	payload := map[string]string{
		"message": "Successful deleted",
	}

	response.AsJson(w, http.StatusOK, payload)
}
