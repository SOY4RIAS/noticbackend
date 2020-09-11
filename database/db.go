package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"noticbackend/config"
	"time"
)

var client *mongo.Client
var err error

func GetClient(c *config.Config) *mongo.Client {

	if client == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err = mongo.Connect(ctx, options.Client().ApplyURI(c.DatabaseUri))

		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(context.TODO(), nil)

		if err != nil {
			log.Fatal(err)
		}



		fmt.Println("Connected to MongoDB!")
	}

	return client
}

func GetCollection(ctx context.Context, opts GetCollectionOptions) (*mongo.Collection,  error) {

	//if err := opts.Client.Connect(ctx); err != nil {
	//	return nil, nil, err
	//}
	//cancel := func() error {
	//	return opts.Client.Disconnect(ctx)
	//}

	err = opts.Client.Ping(ctx, nil)

	if err != nil {
		return nil,  err
	}


	collection := opts.Client.Database(opts.DatabaseName).Collection(opts.CollectionName)


	return collection, nil
}
