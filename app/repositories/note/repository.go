package note

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"noticbackend/app/models"
)

type Repository interface {

	// Create, will perform db opration to save user
	// Returns modified user and error if occurs
	Create(context.Context, *models.Note) error

	// FildAll, returns all users in the system
	// It will return error also if occurs
	FindAll(context.Context) ([]models.Note, error)

	// FindOneById, find the user by the provided id
	// return matched user and error if any
	FindOneById(context.Context, primitive.ObjectID) (*models.Note, error)

	// Update, will update user data by id
	// return error if any
	Update(context.Context, primitive.ObjectID, models.NoteUpdate) error

	// Delete, will remove user entry from DB
	// Return error if any
	Delete(context.Context, primitive.ObjectID) error
}
