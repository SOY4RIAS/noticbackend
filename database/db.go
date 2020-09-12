package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"noticbackend/config"
	"sync"
	"time"
)

var (
	db     *mongo.Database
	client *mongo.Client
	once   sync.Once
	err    error
)

func New(c *config.Config) *mongo.Database {
	once.Do(func() {
		setClient(c)
	})
	return db
}

func setClient(c *config.Config) {

	if client == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err = mongo.Connect(ctx, options.Client().ApplyURI(c.DatabaseUri))

		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(ctx, nil)

		if err != nil {
			log.Fatal(err)
		}

		db = client.Database(c.DatabaseName)

		fmt.Println("Connected to MongoDB!")
	}

}

func GetCollection(ctx context.Context, collectionName string) (*mongo.Collection, error) {

	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	collection := db.Collection(collectionName)

	return collection, nil
}
