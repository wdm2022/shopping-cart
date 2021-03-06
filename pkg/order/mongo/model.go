package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	OrderId   = "_id"
	Paid      = "paid"
	Items     = "items"
	UserId    = "user_id"
	TotalCost = "total_cost"
)

type Order struct {
	OrderId   primitive.ObjectID   `bson:"_id"`
	Paid      bool                 `bson:"paid"`
	Items     []primitive.ObjectID `bson:"items"`
	UserId    primitive.ObjectID   `bson:"user_id"`
	TotalCost int64                `bson:"total_cost"`
}

type Log struct {
	TxId    primitive.ObjectID `bson:"_id"`
	OrderId primitive.ObjectID `bson:"order_id"`
	Time    primitive.DateTime `bson:"time"`
	Status  string             `bson:"status"` //"started,done,reverted"
}
