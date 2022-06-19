package mongo

import (
	"context"
	"log"
	"time"

	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(config *Config) mongoDriver.Client {
	credential := options.Credential{
		Username:   config.Username,
		Password:   config.Password,
		AuthSource: config.Database,
	}

	clientOptions := options.Client()
	clientOptions.SetAuth(credential)
	clientOptions.SetHosts(config.Hosts)
	clientOptions.SetDirect(config.DirectConnection)

	client, err := mongoDriver.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Error when creating MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error when connecting to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error when pinging MongoDB: %v", err)
	}

	log.Printf("Connected to MongoDB")
	return *client
}
