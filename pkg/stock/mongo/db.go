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

type StockConnection struct {
	db.MongoConnection
	StockCollection *mongo.Collection
	//add a context with timeout?
}

func Init(client *mongo.Client) *StockConnection {
	database := client.Database(DbName)
	return &StockConnection{
		MongoConnection: db.MongoConnection{
			Database: database,
			Client:   client,
		},
		StockCollection: database.Collection(CollectionName),
	}
}

/*
Create a new item with the given price and return its id
*/
func (o *StockConnection) NewItem(price int64) (string, error) {

	stock := Stock{
		ItemId: primitive.NewObjectID(),
		Price:  price,
		Amount: 0,
	}

	res, err := o.StockCollection.InsertOne(context.Background(),
		stock)

	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil

}

func (o *StockConnection) addStock(itemId string, amount int64) error {

	objId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return err
	}
	query := bson.D{{ItemId, objId}}

	add := bson.D{{"$inc", bson.D{{StockAmount, amount}}}}

	res, err := o.StockCollection.UpdateOne(context.Background(), query, add)

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

func (o *StockConnection) SubtractStock(itemId string, amount int) error {

	objId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return err
	}
	query := bson.D{{ItemId, objId}}

	add := bson.D{{"$inc", bson.D{{StockAmount, 0 - (amount)}}}}

	res, err := o.StockCollection.UpdateOne(context.Background(), query, add)

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

func (o *StockConnection) findStock(itemId string) (*Stock, error) {

	objId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return nil, err
	}
	query := bson.D{{ItemId, objId}}

	res := o.StockCollection.FindOne(context.Background(), query)

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
