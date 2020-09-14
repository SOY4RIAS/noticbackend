package note

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"noticbackend/app/models"
	"noticbackend/config"
	"noticbackend/database"
	"noticbackend/utils/parse"
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

func (n RepositoryNote) FindOneById(ctx context.Context, id primitive.ObjectID) (*models.Note, error) {

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

func (n RepositoryNote) Update(ctx context.Context, id primitive.ObjectID, noteUpdate models.NoteUpdate) error {
	collection, err := database.GetCollection(ctx, collectionName)

	if err != nil {
		return err
	}

	filter := bson.D{{"_id", id}, {"isDeleted", false}}

	noteUpdate.UpdatedAt = time.Now().Unix()

	updateDocument, err := parse.ToDoc(noteUpdate)

	if err != nil {
		return err
	}

	update := bson.D{{"$set", updateDocument}}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (n RepositoryNote) Delete(ctx context.Context, id primitive.ObjectID) error {
	collection, err := database.GetCollection(ctx, collectionName)

	if err != nil {
		return err
	}

	filter := bson.D{{"_id", id}, {"isDeleted", false}}

	noteDelete := models.NoteDelete{
		IsDeleted: true,
		DeletedAt: time.Now().Unix(),
	}

	deleteDocument, err := parse.ToDoc(noteDelete)

	if err != nil {
		return err
	}

	deleteStmt := bson.D{{"$set", deleteDocument}}

	_, err = collection.UpdateOne(ctx, filter, deleteStmt)

	if err != nil {
		return err
	}

	return nil
}
