package db

import "go.mongodb.org/mongo-driver/mongo"

type MongoConnection struct {
	Database *mongo.Database
	Client   *mongo.Client
}
