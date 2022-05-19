package db

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	OrderId   = "_id"
	Paid      = "paid"
	Items     = "items"
	UserId    = "user_id"
	TotalCost = "total_cost"
)

type Order struct {
	OrderId   primitive.ObjectID `bson:"_id"`
	Paid      bool               `bson:"paid"`
	Items     []string           `bson:"items"`
	UserId    primitive.ObjectID `bson:"user_id"`
	TotalCost float64            `bson:"total_cost"`
}
