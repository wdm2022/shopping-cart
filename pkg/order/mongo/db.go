package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"shopping-cart/pkg/db"
)

const (
	DbName         = "order"
	CollectionName = "orders"
)

type OrdersConnection struct {
	db.MongoConnection
	OrderCollection *mongo.Collection
	//add a context with timeout?
}

func Init(client *mongo.Client) *OrdersConnection {
	database := client.Database(DbName)
	return &OrdersConnection{
		MongoConnection: db.MongoConnection{
			Database: database,
			Client:   client,
		},
		OrderCollection: database.Collection(CollectionName),
	}
}

/*
Deletes an order with a specific id
@param id id in hex format without the '0x' prefix
*/
func (orderConn *OrdersConnection) DeleteOrder(id []byte) error {
	res, err := orderConn.OrderCollection.DeleteOne(context.Background(), bson.D{{OrderId, id}})
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
func (orderConn *OrdersConnection) EmptyOrder(userId []byte) ([]byte, error) {

	order := Order{
		OrderId:   primitive.NewObjectID(),
		Paid:      false,
		Items:     []primitive.ObjectID{},
		UserId:    interface{}(userId).(primitive.ObjectID),
		TotalCost: 0,
	}

	res, err := orderConn.OrderCollection.InsertOne(context.Background(),
		order)
	if err != nil {
		return []byte{}, err
	}
	return interface{}(res.InsertedID).([]byte), nil
}

/*
Find one order by id
the id is in hex format
*/
func (orderConn *OrdersConnection) FindOrder(id []byte) (*Order, error) {

	res := orderConn.OrderCollection.FindOne(context.Background(),
		bson.D{{OrderId, id}})

	if res.Err() != nil {
		return nil, res.Err()
	}
	order := &Order{}
	err := res.Decode(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

/*
Add an item to an order
the ids are in hex format
*/
func (orderConn *OrdersConnection) AddItem(orderId []byte, itemId []byte) error {

	push := bson.D{{"$push", bson.D{{Items, itemId}}}}

	query := bson.D{{OrderId, orderId}}

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

/*
Remove an item from an order
the ids are in hex format
*/
func (orderConn *OrdersConnection) RemoveItem(orderId []byte, itemId []byte) error {
	pull := bson.D{{"$pull", bson.D{{Items, itemId}}}}

	query := bson.D{{OrderId, orderId}}

	res, err := orderConn.OrderCollection.UpdateOne(context.Background(), query, pull)

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
