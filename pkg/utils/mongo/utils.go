package mongo

import (
	"context"
	"fmt"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func Connect(config *Config) mongoDriver.Client {
	credential := options.Credential{
		Username:   config.Username,
		Password:   config.Password,
		AuthSource: config.Database,
	}

	// TODO: Handle multiple hosts when working with Mongo replica set
	hosts := []string{fmt.Sprintf("%s:%v", config.Host, config.Port)}

	clientOptions := options.Client()
	clientOptions.SetAuth(credential)
	clientOptions.SetHosts(hosts)

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
