package note

import (
	"context"
	"noticbackend/app/models"
)

type Service interface {
	//Update(context.Context, string, *models.NoteUpdate) error
	Create(context.Context, *models.Note) error
}