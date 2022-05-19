package db

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

const (
	MongoUriEnvKey = "MONGO_URI"
)

func InitClient() (*mongo.Client, error) {
	uri, ok := os.LookupEnv(MongoUriEnvKey)
	if !ok {
		return nil, errors.New("environment variableN not set")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err // unwrap
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//defer done()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err // unwrap
	}
	return client, nil
}
