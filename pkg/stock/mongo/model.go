package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	ItemId      = "_id"
	StockAmount = "stock"
	Price       = "price"
)

type Stock struct {
	ItemId primitive.ObjectID `bson:"_id"`
	Price  int64              `bson:"price"`
	Amount int64              `bson:"stock"`
}
