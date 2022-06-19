package mongo

import (
	"context"
	"errors"
	"fmt"
	sf "github.com/sa-/slicefunk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"shopping-cart/pkg/db"
	"shopping-cart/pkg/utils"
)

const (
	DbName              = "stock"
	StockCollectionName = "stock"
	LogCollectionName   = "logs"
)

type StockConnection struct {
	db.MongoConnection
	StockCollection *mongo.Collection
	LogCollection   *mongo.Collection
	//add a context with timeout?
}

func Init(client *mongo.Client) *StockConnection {
	database := client.Database(DbName)
	return &StockConnection{
		MongoConnection: db.MongoConnection{
			Database: database,
			Client:   client,
		},
		StockCollection: database.Collection(StockCollectionName),
		LogCollection:   database.Collection(LogCollectionName),
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

func (o *StockConnection) AddStock(itemId string, amount int64) error {

	objId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return err
	}
	query := bson.D{
		primitive.E{
			Key:   ItemId,
			Value: objId,
		},
	}

	add := bson.D{
		primitive.E{
			Key: "$inc",
			Value: bson.D{
				primitive.E{
					Key:   StockAmount,
					Value: amount,
				},
			},
		},
	}

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

func (o *StockConnection) SubtractStock(itemId string, amount int64) error {

	objId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return err
	}
	query := bson.D{
		primitive.E{
			Key:   ItemId,
			Value: objId,
		},
	}

	add := bson.D{
		primitive.E{
			Key: "$inc",
			Value: bson.D{
				primitive.E{
					Key:   StockAmount,
					Value: 0 - (amount),
				},
			},
		},
	}

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

func (o *StockConnection) FindStock(itemId string) (*Stock, error) {

	objId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return nil, err
	}
	query := bson.D{
		primitive.E{
			Key:   ItemId,
			Value: objId,
		},
	}

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

func (o *StockConnection) CalculateTotalCost(itemIds []string) (int64, error) {
	objIds := sf.Map(itemIds, func(t string) primitive.ObjectID {
		id, err := primitive.ObjectIDFromHex(t)
		if err != nil {
			return [12]byte{}
		}
		return id
	})
	ctx, cancelFunc := utils.ContextWithTimeOut()
	defer cancelFunc()
	aggregate, err := o.StockCollection.Aggregate(ctx, []bson.M{
		{
			"$match": bson.M{
				ItemId: bson.M{
					"$in": objIds,
				},
			},
		},
		{
			"$group": bson.M{
				"_id": nil,
				"total": bson.M{
					"$sum": "$price",
				},
			},
		}, /*, {
			"$project": bson.M{
				"total": bson.M{
					"$toString": "$total1",
				},
			},
		},*/
	})
	if err != nil {
		return 0, err
	}

	//idk mongo gives this format as a result : {"_id": null,"total": {"$numberLong":"1"}} if I don't convert to string

	// screw this, just do it manually
	//type costRes struct {
	//	id    primitive.ObjectID `bson:"_id"`
	//	total string             `bson:"total"`
	//}

	aggregate.Next(ctx)
	totcalCost, ok := aggregate.Current.Index(1).Value().Int64OK()
	if !ok {
		return 0, errors.New("failed to convert to int")
	}
	if aggregate.Current == nil {
		// no result
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return totcalCost, nil
}

func (o *StockConnection) SubtractBatchStock(txId string, itemIds []string) error {
	// TODO: Add DB call to remove the stock using the provided list of items. #Rahim :)

	objTxId, err := primitive.ObjectIDFromHex(txId)
	if err != nil {
		return err
	}

	amounts := make(map[primitive.ObjectID]int64)
	objIds := sf.Map(itemIds, func(t string) primitive.ObjectID {
		objId, _ := primitive.ObjectIDFromHex(t)
		amounts[objId] = amounts[objId] + 1
		return objId
	})
	ctx, cancel := utils.ContextWithTimeOut()
	defer cancel()
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {

		logRes := o.LogCollection.FindOne(sessCtx, bson.D{primitive.E{Key: "_id", Value: objTxId}})

		if logRes.Err() == mongo.ErrNoDocuments {
			// we have not handled this yet
		} else if logRes.Err() != nil {
			return false, logRes.Err()
		} else {
			//we have handled this
			return true, nil
		}

		//todo How do to all of this in one call to database.
		for _, id := range objIds {
			query := bson.M{"_id": id, StockAmount: bson.M{"$gte": amounts[id]}}
			update := bson.M{"$inc": bson.M{StockAmount: -amounts[id]}}

			res := o.StockCollection.FindOneAndUpdate(sessCtx, query, update)
			if res.Err() != nil {
				return nil, res.Err()
			}
		}

		_, err = o.LogCollection.InsertOne(sessCtx, Log{
			TxId:   objTxId,
			Status: "done",
			Items:  objIds,
		})
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
	session, err := o.Client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}

	return nil
}

func (o *StockConnection) RollBack(txId string) error {
	// TODO: Add DB call to remove the stock using the provided list of items. #Rahim :)

	objTxId, err := primitive.ObjectIDFromHex(txId)
	if err != nil {
		return err
	}

	ctx, cancel := utils.ContextWithTimeOut()
	defer cancel()
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {

		logRes := o.LogCollection.FindOneAndUpdate(sessCtx,
			bson.M{
				"_id": objTxId,
				"status": bson.M{
					"$ne": "reverted",
				},
			},
			bson.M{
				"$set": bson.M{
					"status": "reverted",
				},
			})

		if logRes.Err() == mongo.ErrNoDocuments {
			// we have not handled this yet
			fmt.Println("allready handled ", txId)
			return nil, nil
		} else if logRes.Err() != nil {

			return false, logRes.Err()
		}
		log := &Log{}

		err := logRes.Decode(log)
		if err != nil {
			return nil, err
		}

		amounts := make(map[primitive.ObjectID]int64)
		_ = sf.Map(log.Items, func(t primitive.ObjectID) error {
			amounts[t] = amounts[t] + 1
			return nil
		})

		//todo How do to all of this in one call to database.
		for _, id := range log.Items {
			query := bson.M{"_id": id}
			update := bson.M{"$inc": bson.M{StockAmount: amounts[id]}}

			res := o.StockCollection.FindOneAndUpdate(sessCtx, query, update)
			if res.Err() != nil {
				return nil, res.Err()
			}
		}

		return nil, nil
	}
	session, err := o.Client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}

	return nil
}

func (o *StockConnection) AddBatchStock(itemIds []string) error {
	// TODO: Add DB call to add the stock using the provided list of items. #Rahim :)
	return nil
}
