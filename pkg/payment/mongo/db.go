package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"
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
		Orders: &[]Order{},
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
	objUserId, err := primitive.ObjectIDFromHex(userId)
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
				log.Println("Error when finding log ", logRes.Err())
				return false, logRes.Err()
			} else {
				//we have handled this
				return true, nil
			}
		}
		// find user that doesn't have this order, so we haven't paid it yet
		query := primitive.M{
			UserId: objUserId,
			Orders: primitive.M{
				"$not": primitive.M{
					"$all": primitive.A{
						primitive.M{
							OrderId:   objOrderId,
							OrderPaid: true,
						},
					},
				},
			},
			Credit: bson.M{
				"$gte": amount,
			},
		}
		decFunc := bson.M{
			"$push": primitive.M{
				Orders: primitive.M{
					OrderId:   objOrderId,
					OrderPaid: true,
				},
			},
			"$inc": primitive.M{
				Credit: 0 - amount,
			},
		}
		res := p.PaymentCollection.FindOneAndUpdate(sessCtx, query, decFunc)
		if res.Err() != nil {
			log.Println("error when finding user and paying ", userId, " error ", res.Err())
			return false, res.Err()
		}

		if txId != emptyHex {
			_, err := p.LogCollection.InsertOne(sessCtx, Log{
				TxId:    objTxId,
				Status:  "done",
				Amount:  amount,
				OrderId: objOrderId,
				UserId:  objUserId,
			})

			if err != nil {
				return false, err
			}
		}
		return true, nil
	}

	session, err := p.Client.StartSession()
	if err != nil {
		log.Println("error when starting session ", err)
		return false, err
	}
	defer session.EndSession(ctx)

	result, err := session.WithTransaction(ctx, callBack)

	return result.(bool), err
}

func (p *PaymentConnection) Rollback(txId string) error {
	objTxId, err := primitive.ObjectIDFromHex(txId)
	if err != nil {
		return err
	}

	ctx, cancel := utils.ContextWithTimeOut()
	defer cancel()
	callBack := func(sessCtx mongo.SessionContext) (interface{}, error) {

		query := bson.M{
			LogId: objTxId,
			Status: bson.M{
				"$ne": "reverted",
			},
		}
		update := bson.M{
			"$set": bson.M{
				Status: "reverted",
			},
		}
		logRes := p.LogCollection.FindOneAndUpdate(sessCtx, query, update)
		if logRes.Err() != nil {
			return nil, logRes.Err()
		}
		payLog := &Log{}
		err := logRes.Decode(payLog)
		if err != nil {
			return nil, err
		}

		query = bson.M{
			"_id": payLog.UserId,
		}
		update = bson.M{
			"$inc": bson.M{
				Credit: payLog.Amount,
			},
			"$pull": bson.M{
				"orders": bson.M{"order_id": payLog.OrderId},
			},
		}

		userRes := p.PaymentCollection.FindOneAndUpdate(sessCtx, query, update)
		if userRes.Err() != nil {
			fmt.Println("failed to rollbackk tx :  ", payLog.TxId, " user ", payLog.UserId)
			return nil, userRes.Err()
		}

		return nil, nil
	}
	session, err := p.Client.StartSession()
	if err != nil {
		log.Println("error when starting session ", err)
		return err
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, callBack)

	return err
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
