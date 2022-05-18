package mongo

import mongoDriver "go.mongodb.org/mongo-driver/mongo"
import mongoUtils "shopping-cart/pkg/utils/mongo"

var (
	client mongoDriver.Client
)

func Connect(options *mongoUtils.Config) mongoDriver.Client {
	client = mongoUtils.Connect(options)
	return client
}
