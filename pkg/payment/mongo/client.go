package mongo

import (
	mongoUtils "shopping-cart/pkg/utils/mongo"

	mongoDriver "go.mongodb.org/mongo-driver/mongo"
)

var (
	client   mongoDriver.Client
	database string
)

func Connect(options *mongoUtils.Config) mongoDriver.Client {
	client = mongoUtils.Connect(options)
	database = options.Database
	return client
}
