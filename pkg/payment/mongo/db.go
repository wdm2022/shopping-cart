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
	DbName         = "payment"
	CollectionName = "payments"
)

type PaymentConnection struct {
	db.MongoConnection
	PaymentCollection *mongo.Collection
	//add a context with timeout?
}

func Init(client *mongo.Client) *PaymentConnection {
	database := client.Database(DbName)
	return &PaymentConnection{
		MongoConnection: db.MongoConnection{
			Database: database,
			Client:   client,
		},
		PaymentCollection: database.Collection(CollectionName),
	}
}

/*
create and return id of created user
*/
func (p *PaymentConnection) CreateUser() (string, error) {
	user := User{
		UserId: primitive.NewObjectID(),
		Credit: 0,
	}

	res, err := p.PaymentCollection.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (p *PaymentConnection) FindUser(userId string) (*User, error) {
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	res := p.PaymentCollection.FindOne(context.Background(),
		bson.D{{UserId, objId}})

	if res.Err() != nil {
		return nil, res.Err()
	}
	user := &User{}
	err = res.Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p *PaymentConnection) AddFunds(userId string, amount int64) error {
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil
	}

	query := bson.D{{UserId, objId}}

	add := bson.D{{"$inc", bson.D{{Credit, amount}}}}

	res, err := p.PaymentCollection.UpdateOne(context.Background(), query, add)

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
