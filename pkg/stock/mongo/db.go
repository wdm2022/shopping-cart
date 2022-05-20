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
	DbName         = "stock"
	CollectionName = "stock"
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
Create a new item with the given price and return its id
*/
func (o *OrdersConnection) NewItem(price float64) (string, error) {

	stock := Stock{
		ItemId: primitive.NewObjectID(),
		Price:  price,
		Amount: 0,
	}

	res, err := o.OrderCollection.InsertOne(context.Background(),
		stock)

	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil

}

func (o *OrdersConnection) addStock(itemId string, amount uint64) error {

	objId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return err
	}
	query := bson.D{{ItemId, objId}}

	add := bson.D{{"$inc", bson.D{{StockAmount, amount}}}}

	res, err := o.OrderCollection.UpdateOne(context.Background(), query, add)

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

func (o *OrdersConnection) SubtractStock(itemId string, amount uint64) error {

	objId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return err
	}
	query := bson.D{{ItemId, objId}}

	add := bson.D{{"$inc", bson.D{{StockAmount, 0 - (int64(amount))}}}}

	res, err := o.OrderCollection.UpdateOne(context.Background(), query, add)

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

func (o *OrdersConnection) findStock(itemId string) (*Stock, error) {

	objId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return nil, err
	}
	query := bson.D{{ItemId, objId}}

	res := o.OrderCollection.FindOne(context.Background(), query)

	if res.Err() != nil {
		return nil, res.Err()
	}

	stock := &Stock{}

	err = res.Decode(stock)
	if err != nil {
		return nil, err
	}
	return stock, nil
}
