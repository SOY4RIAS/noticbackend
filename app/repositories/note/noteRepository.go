package note

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"noticbackend/app/models"
	"noticbackend/config"
	"noticbackend/database"
)

const collectionName = "notes"

type NoteRepository struct {
	client *mongo.Client
	config *config.Config
}

func (NoteRepository) New(client *mongo.Client, config *config.Config) *NoteRepository {
	return &NoteRepository{client: client, config: config}
}

func (n *NoteRepository) Create(ctx context.Context, note *models.Note) error {

	opts := database.GetCollectionOptions{
		Client: n.client,
		DatabaseName: n.config.DatabaseName,
		CollectionName: collectionName,
	}

	collection, err := database.GetCollection(ctx, opts)
	fmt.Println(collection)

	if err != nil {
		return err
	}



	result, err := collection.InsertOne(ctx, note)

	if err != nil {
		return err
	}

	fmt.Println(result.InsertedID)

	return nil
}

func (n NoteRepository) FindAll(ctx context.Context) ([]*models.Note, error) {
	panic("implement me")
}

func (n NoteRepository) FindOneById(ctx context.Context, s string) (*models.Note, error) {
	panic("implement me")
}

func (n NoteRepository) Update(ctx context.Context, i interface{}, i2 interface{}) error {
	panic("implement me")
}

func (n NoteRepository) Delete(ctx context.Context, note *models.Note) error {
	panic("implement me")
}
