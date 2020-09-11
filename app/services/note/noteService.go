package note

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"noticbackend/app/models"
	"noticbackend/app/repositories/note"
	"noticbackend/config"
)

type NoteService struct {
	client *mongo.Client
	repository note.Repository
	config *config.Config
}
func (NoteService) New(client *mongo.Client, config *config.Config) *NoteService {
	repository := note.NoteRepository{}.New(client,config)

	return &NoteService{client: client, repository: repository, config: config}
}

func (n *NoteService) Create(ctx context.Context, note *models.Note) error {
	return n.repository.Create(ctx, note)
}

