package note

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"noticbackend/app/models"
	"noticbackend/app/repositories/note"
	"noticbackend/config"
)

// ServiceNote service of the notes
type ServiceNote struct {
	db         *mongo.Database
	repository note.Repository
	config     *config.Config
}

func (ServiceNote) New(db *mongo.Database, config *config.Config) *ServiceNote {
	repository := note.RepositoryNote{}.New(db, config)

	return &ServiceNote{db: db, repository: repository, config: config}
}

func (n *ServiceNote) Create(ctx context.Context, note *models.Note) error {
	return n.repository.Create(ctx, note)
}

func (n *ServiceNote) FindAll(ctx context.Context) ([]models.Note, error) {
	return n.repository.FindAll(ctx)
}

func (n *ServiceNote) FindOneById(ctx context.Context, id string) (*models.Note, error) {
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	return n.repository.FindOneById(ctx, objID)
}

func (n *ServiceNote) Update(ctx context.Context, id string, noteUpdate models.NoteUpdate) error {
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	return n.repository.Update(ctx, objID, noteUpdate)
}

func (n *ServiceNote) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	return n.repository.Delete(ctx, objID)
}
