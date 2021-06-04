package data

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

// MongoWordList is a high level representation of the words database
// implemented in MongoDB.
type MongoWordList struct {
	client *mongo.Client
	words  *mongo.Collection
}

// ConnectMongo connects to the Mongo database with the URI and return
// a constructed WordList.
func ConnectMongo(ctx context.Context, uri string) (*MongoWordList, error) {
	cs, err := connstring.ParseAndValidate(uri)
	if err != nil {
		return nil, fmt.Errorf("invalid uri: %s: %w", uri, err)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo database (uri: %s): %w", uri, err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("mongo client does not respond: %w", err)
	}

	db := client.Database(cs.Database)
	col := db.Collection("words")

	return &MongoWordList{
		client: client,
		words:  col,
	}, nil
}

// Word generates a random word of length l from the wordlist.
func (wl *MongoWordList) Word(ctx context.Context, l int16) (string, error) {
	w, err := wl.randomWord(ctx, l)
	if err != nil {
		return "", fmt.Errorf("failed to generate a random word: %w", err)
	}

	return w, nil
}

// Close closes all idle connection pools.
func (wl *MongoWordList) Close(ctx context.Context) error {
	if err := wl.client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from Mongo database: %w", err)
	}

	return nil
}
