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
func (orderConn *OrdersConnection) DeleteOrder(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	res, err := orderConn.OrderCollection.DeleteOne(context.Background(), bson.D{{OrderId, objId}})
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
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res := orderConn.OrderCollection.FindOne(context.Background(),
		bson.D{{OrderId, objId}})

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
		return nil
	}

	objItemId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return nil
	}

	push := bson.D{{"$push", bson.D{{Items, objItemId}}}}

	query := bson.D{{OrderId, objOrderId}}

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
func (orderConn *OrdersConnection) RemoveItem(orderId string, itemId string) error {
	objOrderId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return nil
	}

	objItemId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return nil
	}

	pull := bson.D{{"$pull", bson.D{{Items, objItemId}}}}

	query := bson.D{{OrderId, objOrderId}}

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
