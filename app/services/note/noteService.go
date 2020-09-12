package note

import (
	"context"
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
