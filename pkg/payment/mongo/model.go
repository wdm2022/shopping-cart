package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	UserId    = "_id"
	Credit    = "credit"
	OrderPaid = "paid"
	OrderId   = "order_id"
)

type Order struct {
	OrderId primitive.ObjectID `bson:"order_id"`
	Paid    bool               `bson:"paid"`
}

type User struct {
	UserId primitive.ObjectID `bson:"_id"`
	Credit int64              `bson:"credit"`
}
