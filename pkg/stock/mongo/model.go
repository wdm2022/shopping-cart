package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	ItemId      = "_id"
	StockAmount = "stock"
	Price       = "price"
)

type Stock struct {
	ItemId primitive.ObjectID `bson:"_id"`
	Price  float64            `bson:"price"`
	Amount uint64             `bson:"stock"`
}
