package mongo

import (
	"context"
	"errors"
	"fmt"
	"shopping-cart/pkg/db"
	"shopping-cart/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DbName              = "order"
	OrderCollectionName = "orders"
	LogCollectionName   = "log"
)

type OrdersConnection struct {
	db.MongoConnection
	OrderCollection *mongo.Collection
	LogCollection   *mongo.Collection
	//ctx             context.Context
}

func Init(client *mongo.Client) *OrdersConnection {
	database := client.Database(DbName)
	//todo cancel fun
	return &OrdersConnection{
		MongoConnection: db.MongoConnection{
			Database: database,
			Client:   client,
		},
		OrderCollection: database.Collection(OrderCollectionName),
		LogCollection:   database.Collection(LogCollectionName),
	}
}

/*
Deletes an order with a specific id
@param id id in hex format without the '0x' prefix
*/
func (orderConn *OrdersConnection) DeleteOrder(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	query := bson.D{
		primitive.E{
			Key:   OrderId,
			Value: objId,
		},
	}
	res, err := orderConn.OrderCollection.DeleteOne(context.Background(), query)
	if err != nil {
		return err
	}
	if res.DeletedCount != 1 {
		return errors.New("document was not deleted")
	}
	return nil
}

/*
Create a new empty order for the user
The id is in hex format
*/
func (orderConn *OrdersConnection) EmptyOrder(userId string) (string, error) {
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return "", err
	}

	order := Order{
		OrderId:   primitive.NewObjectID(),
		Paid:      false,
		Items:     []primitive.ObjectID{},
		UserId:    objId,
		TotalCost: 0,
	}

	res, err := orderConn.OrderCollection.InsertOne(context.Background(),
		order)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

/*
Find one order by id
the id is in hex format
*/
func (orderConn *OrdersConnection) FindOrder(id string) (*Order, error) {
	// TODO: Remove totalcost from this table, is retrieved from the stock service #Rahim
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	query := bson.D{
		primitive.E{
			Key:   OrderId,
			Value: objId,
		},
	}
	res := orderConn.OrderCollection.FindOne(context.Background(), query)

	if res.Err() != nil {
		return nil, res.Err()
	}
	order := &Order{}
	err = res.Decode(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

/*
Add an item to an order
the ids are in hex format
*/
func (orderConn *OrdersConnection) AddItem(orderId string, itemId string) error {
	objOrderId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return err
	}

	objItemId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return err
	}

	push := bson.D{
		primitive.E{
			Key: "$push",
			Value: bson.D{
				primitive.E{
					Key:   Items,
					Value: objItemId,
				},
			},
		},
	}

	query := bson.D{
		primitive.E{
			Key:   OrderId,
			Value: objOrderId,
		},
	}

	res, err := orderConn.OrderCollection.UpdateOne(context.Background(), query, push)

	if err != nil {
		return err
	}
	if res.ModifiedCount > 1 {
		return errors.New("updated multiple documents")
	}
	if res.ModifiedCount == 0 {
		return errors.New("updated 0 documents")
	}
	return nil
}

func (orderConn *OrdersConnection) PayOrder(orderId string) error {
	objOrderId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"paid": true,
		},
	}

	query := bson.D{
		primitive.E{
			Key:   OrderId,
			Value: objOrderId,
		},
	}

	res, err := orderConn.OrderCollection.UpdateOne(context.Background(), query, update)

	if err != nil {
		return err
	}
	if res.ModifiedCount > 1 {
		return errors.New("updated multiple documents")
	}
	if res.ModifiedCount == 0 {
		return errors.New("updated 0 documents")
	}
	return nil
}

func (orderConn *OrdersConnection) UnpayOrder(orderId string) error {
	objOrderId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"paid": false,
		},
	}

	query := bson.M{
		OrderId: objOrderId,
	}

	res, err := orderConn.OrderCollection.UpdateOne(context.Background(), query, update)

	if err != nil {
		return err
	}
	if res.ModifiedCount > 1 {
		return errors.New("updated multiple documents")
	}
	if res.ModifiedCount == 0 {
		fmt.Println("updated 0")
	}
	return nil
}

func (orderConn *OrdersConnection) StartTransaction(txId string, orderId string) error {
	objTxId, err := primitive.ObjectIDFromHex(txId)
	if err != nil {
		return err
	}
	objOrderId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return err
	}
	ctx, cancel := utils.ContextWithTimeOut()
	defer cancel()

	_, err = orderConn.LogCollection.InsertOne(ctx, Log{
		TxId:    objTxId,
		Time:    primitive.NewDateTimeFromTime(time.Now()),
		OrderId: objOrderId,
		Status:  "started",
	})
	if err != nil {
		return err
	}

	return nil

}

func (orderConn *OrdersConnection) EndTransaction(txId string) error {
	objTxId, err := primitive.ObjectIDFromHex(txId)
	if err != nil {
		return err
	}

	ctx, cancel := utils.ContextWithTimeOut()
	defer cancel()

	query := primitive.M{
		"_id": objTxId,
	}

	//todo don't use literals
	update := primitive.M{
		"$set": primitive.M{
			"status": "done",
		}}

	_, err = orderConn.LogCollection.UpdateOne(ctx, query, update)
	if err != nil {
		return err
	}

	return nil

}

func (orderConn *OrdersConnection) Lock(txId string) error {
	objTxId, err := primitive.ObjectIDFromHex(txId)
	if err != nil {
		return err
	}

	ctx, cancel := utils.ContextWithTimeOut()
	defer cancel()

	query := primitive.M{
		"_id": objTxId,
	}

	//todo don't use literals
	update := primitive.M{
		"$set": primitive.M{
			"time": primitive.NewDateTimeFromTime(time.Now()),
		}}

	res := orderConn.LogCollection.FindOneAndUpdate(ctx, query, update)
	if res.Err() != nil {
		return res.Err()
	}

	return nil

}

func (orderConn *OrdersConnection) RevertTransaction(txId string) error {
	objTxId, err := primitive.ObjectIDFromHex(txId)
	if err != nil {
		return err
	}

	ctx, cancel := utils.ContextWithTimeOut()
	defer cancel()

	query := primitive.M{
		"_id": objTxId,
	}

	//todo don't use literals
	update := primitive.M{
		"$set": primitive.M{
			"status": "reverted",
		}}

	_, err = orderConn.LogCollection.UpdateOne(ctx, query, update)
	if err != nil {
		return err
	}

	return nil

}

func (orderConn *OrdersConnection) FindOpenTransactions() ([]Log, error) {

	ctx, cancel := utils.ContextWithTimeOut()
	defer cancel()
	query := bson.M{
		"status": "started",
		"time": bson.M{
			"$lt": primitive.NewDateTimeFromTime(time.Now().Add(-time.Second * 20)), // assume that after 20 secs tx is dead
		},
	}

	res, _ := orderConn.LogCollection.Find(ctx, query)

	if res.Err() != nil {
		return nil, res.Err()
	}
	var results []Log

	err := res.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

/*
Remove an item from an order
the ids are in hex format
*/
func (orderConn *OrdersConnection) RemoveItem(orderId string, itemId string) error {
	objOrderId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return err
	}

	objItemId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return err
	}
	//todo, only remove one item of the same id
	pull := bson.D{
		primitive.E{
			Key: "$pull",
			Value: bson.D{
				primitive.E{
					Key:   Items,
					Value: objItemId,
				},
			},
		},
	}
	//pull := bson.D{{"$unset", bson.D{{"items.$[]", objItemId}}}}

	query := bson.D{
		primitive.E{
			Key:   OrderId,
			Value: objOrderId,
		},
	}

	res, err := orderConn.OrderCollection.UpdateOne(context.Background(), query, pull)

	if err != nil {
		return err
	}
	if res.ModifiedCount > 1 {
		return errors.New("updated multiple documents")
	}
	//todo think about this, I guess we probably want to delete if it is there and not error when it isn't there
	//if res.ModifiedCount == 0 {
	//	return errors.New("updated 0 documents")
	//}
	return nil
}
