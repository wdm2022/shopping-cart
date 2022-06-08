package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	ItemId      = "_id"
	StockAmount = "stock"
	Price       = "price"
)

type Log struct {
	TxId   primitive.ObjectID   `bson:"_id"`
	Status string               `bson:"status"`
	Items  []primitive.ObjectID `bson:"items"`
}

type Stock struct {
	ItemId primitive.ObjectID `bson:"_id"`
	Price  int64              `bson:"price"`
	Amount int64              `bson:"stock"`
}
