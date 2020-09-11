package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Note struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title string `json:"name,omitempty" bson:"title,omitempty"`
	Content string `json:"content,omitempty" bson:"content,omitempty"`
	IsDeleted bool `json:"isDeleted,omitempty" bson:"isDeleted,omitempty"`
	CreatedAt int64 `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt int64 `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type NoteUpdate struct {
	Title string `json:"name,omitempty" bson:"title,omitempty"`
	Content string `json:"content,omitempty" bson:"content,omitempty"`
	IsDeleted bool `json:"isDeleted,omitempty" bson:"content,omitempty"`
	UpdatedAt int64 `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

func (note *Note) Initialize()  {
	timestamp := time.Now().Unix()

	note.IsDeleted = false
	note.CreatedAt = timestamp
	note.UpdatedAt = timestamp
}

func (Note) Create() *Note {
	timestamp := time.Now().Unix()

	return &Note{
		IsDeleted: false,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}
}


