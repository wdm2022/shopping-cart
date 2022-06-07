package mongo

import (
	"context"
	"errors"
	"shopping-cart/pkg/db"
	"shopping-cart/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DbName                = "payment"
	PaymentCollectionName = "payments"
	LogCollectionName     = "logs"
)

type PaymentConnection struct {
	db.MongoConnection
	PaymentCollection *mongo.Collection
	LogCollection     *mongo.Collection
	//add a context with timeout?
}

func Init(client *mongo.Client) *PaymentConnection {
	database := client.Database(DbName)
	return &PaymentConnection{
		MongoConnection: db.MongoConnection{
			Database: database,
			Client:   client,
		},
		PaymentCollection: database.Collection(PaymentCollectionName),
		LogCollection:     database.Collection(LogCollectionName),
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
		bson.D{
			primitive.E{
				Key:   UserId,
				Value: objId,
			},
		},
	)

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

	query := bson.D{primitive.E{Key: UserId, Value: objId}}

	add := bson.D{
		primitive.E{
			Key: "$inc",
			Value: bson.D{
				primitive.E{
					Key:   Credit,
					Value: amount,
				},
			},
		},
	}

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

func (p *PaymentConnection) PayOrder(txId string, userId string, orderId string, amount int64) (bool, error) {

	objTxId, err := primitive.ObjectIDFromHex(txId)
	if err != nil {
		return false, err
	}

	objOrderId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return false, err
	}
	objUserId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return false, err
	}

	ctx, cancel := utils.ContextWithTimeOut()
	defer cancel()

	callBack := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// for if this is not used in a sage, but with the "endpoint"
		emptyHex := primitive.ObjectID{}.Hex()
		if txId != emptyHex {
			logRes := p.LogCollection.FindOne(sessCtx, bson.D{{LogId, objTxId}})

			if logRes.Err() == mongo.ErrNoDocuments {
				// we have not handled this yet
			} else if logRes.Err() != nil {
				return false, logRes.Err()
			} else {
				//we have handled this
				return true, nil
			}
		}
		// find user that doesn't have this order, so we haven't paid it yet
		query := bson.D{
			{
				userId, objUserId,
			},
			{Key: Orders, Value: primitive.E{
				Key: "$not",
				Value: primitive.E{
					Key: "$eq",
					Value: primitive.E{
						Key:   OrderId,
						Value: objOrderId,
					},
				},
			}},
		}
		res := p.PaymentCollection.FindOne(sessCtx, query)
		if res.Err() != nil {
			return false, res.Err()
		}
		user := &User{}
		err = res.Decode(user)
		if err != nil {
			return false, err
		}
		newAmount := user.Credit - amount
		if newAmount < 0 {
			return false, errors.New("not sufficient credit")
		}
		decFunc := bson.D{
			primitive.E{Key: "$push", Value: primitive.E{
				Key:   Orders,
				Value: objOrderId,
			}},
			primitive.E{
				Key: "$inc",
				Value: bson.D{
					primitive.E{
						Key:   Credit,
						Value: 0 - (amount),
					},
				},
			},
		}

		if txId != emptyHex {
			_, err := p.LogCollection.InsertOne(sessCtx, Log{
				TxId:    objTxId,
				Status:  "done",
				amount:  amount,
				orderId: objOrderId,
			})

			if err != nil {
				return false, err
			}
		}

		updateRes, err := p.PaymentCollection.UpdateOne(sessCtx, query, decFunc)

		//todo should we check for all of this?
		if err != nil {
			return false, err
		}
		if updateRes.ModifiedCount > 1 {
			return false, errors.New("updated multiple documents")
		}
		if updateRes.ModifiedCount == 0 {
			return false, errors.New("updated 0 documents")
		}
		return true, nil
	}

	session, err := p.Client.StartSession()
	if err != nil {
		return false, err
	}
	defer session.EndSession(ctx)

	result, err := session.WithTransaction(ctx, callBack)

	return result.(bool), err
}

func (p *PaymentConnection) CancelOrder(userId string, orderId string) error {
	objId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return err
	}
	query := bson.D{
		primitive.E{
			Key:   UserId,
			Value: objId,
		},
		primitive.E{
			Key:   OrderId,
			Value: orderId,
		},
	}
	res, err := p.PaymentCollection.DeleteOne(context.Background(), query)
	if err != nil {
		return err
	}
	if res.DeletedCount != 1 {
		return errors.New("document was not deleted")
	}
	return nil
}

func (p *PaymentConnection) StatusPayment(userId string, orderId string) (bool, error) {
	objId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return false, err
	}
	query := bson.D{
		primitive.E{
			Key:   UserId,
			Value: objId,
		},
		primitive.E{
			Key:   OrderId,
			Value: orderId,
		},
	}
	res := p.PaymentCollection.FindOne(context.Background(), query)
	if res.Err() != nil {
		return false, res.Err()
	}
	order := &Order{}
	err = res.Decode(order)
	if err != nil {
		return false, err
	}
	return order.Paid, nil
}
