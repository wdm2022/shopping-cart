package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	UserId    = "_id"
	Credit    = "credit"
	OrderPaid = "paid"
	OrderId   = "order_id"
	Orders    = "orders"
	LogId     = "_id"
	Status    = "status"
	amount    = "amount"
)

type Log struct {
	TxId    primitive.ObjectID `bson:"_id"`
	Status  string             `bson:"status"` // done, reverted
	Amount  int64              `bson:"amount"`
	OrderId primitive.ObjectID `bson:"order_id"`
	UserId  primitive.ObjectID `bson:"user_id"`
}

type Order struct {
	OrderId primitive.ObjectID `bson:"order_id"`
	Paid    bool               `bson:"paid"`
}

type User struct {
	UserId primitive.ObjectID `bson:"_id"`
	Credit int64              `bson:"credit"`
	Orders []Order            `bson:"orders"`
}
