package database

import "go.mongodb.org/mongo-driver/mongo"

type GetCollectionOptions struct {
	Client         *mongo.Client
	DatabaseName   string
	CollectionName string
}
