package note

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"noticbackend/app/models"
	"noticbackend/config"
	"noticbackend/database"
)

const collectionName = "notes"

type RepositoryNote struct {
	db     *mongo.Database
	config *config.Config
}

func (RepositoryNote) New(db *mongo.Database, config *config.Config) *RepositoryNote {
	return &RepositoryNote{db: db, config: config}
}

func (n *RepositoryNote) Create(ctx context.Context, note *models.Note) error {

	collection, err := database.GetCollection(ctx, collectionName)

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

func (n RepositoryNote) FindAll(ctx context.Context) ([]models.Note, error) {
	collection, err := database.GetCollection(ctx, collectionName)

	if err != nil {
		return nil, err
	}

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	var notes []models.Note
	for cur.Next(ctx) {
		var note models.Note

		if err := cur.Decode(&note); err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, nil
}

func (n RepositoryNote) FindOneById(ctx context.Context, id string) (*models.Note, error) {

	collection, err := database.GetCollection(ctx, collectionName)

	if err != nil {
		return nil, err
	}

	var note models.Note

	filter := bson.D{{"_id", id}}

	err = collection.FindOne(context.Background(), filter).Decode(&note)

	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (n RepositoryNote) Update(_ context.Context, _ interface{}, _ interface{}) error {
	panic("implement me")
}

func (n RepositoryNote) Delete(_ context.Context, _ *models.Note) error {
	panic("implement me")
}
