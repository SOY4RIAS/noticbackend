package note

import "noticbackend/app/models"

type FindAll struct {
	Data DataNotes `json:"data"`
}
type FindOneById struct {
	Data DataNote `json:"data"`
}
type DataNote struct {
	Note models.Note `json:"note"`
}
type DataNotes struct {
	Notes []models.Note `json:"notes"`
}
