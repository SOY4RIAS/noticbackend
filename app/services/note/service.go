package note

import (
	"context"
	"noticbackend/app/models"
)

type Service interface {
	//Update(context.Context, string, *models.NoteUpdate) error
	Create(context.Context, *models.Note) error
	FindAll(context.Context) ([]models.Note, error)
	FindOneById(context.Context, string) (*models.Note, error)
	Update(context.Context, string, models.NoteUpdate) error
	Delete(context.Context, string) error
}
