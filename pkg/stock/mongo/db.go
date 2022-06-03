package mongo

import (
	"context"
	"errors"
	sf "github.com/sa-/slicefunk"
	"shopping-cart/pkg/db"
	"shopping-cart/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		bson.M{
			"$match": bson.M{
				ItemId: bson.M{
					"$in": objIds,
				},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": nil,
				"total": bson.M{
					"$sum": "$price",
				},
			},
		},
	})
	if err != nil {
		return 0, err
	}

	type costRes struct {
		id    primitive.ObjectID `bson:"_id"`
		total int64              `bson:"total"`
	}

	totalStruct := &costRes{}

	err = aggregate.Decode(totalStruct)
	if err != nil {
		return 0, err
	}
	return totalStruct.total, nil
}

func (o *StockConnection) SubtractBatchStock(itemIds []string) error {
	// TODO: Add DB call to remove the stock using the provided list of items. #Rahim :)

	amounts := make(map[primitive.ObjectID]int64)
	objIds := sf.Map(itemIds, func(t string) primitive.ObjectID {
		objId, _ := primitive.ObjectIDFromHex(t)
		amounts[objId] = amounts[objId] + 1
		return objId
	})
	ctx, cancel := utils.ContextWithTimeOut()
	defer cancel()
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		for _, id := range objIds {
			query := bson.M{"_id": id}
			add := bson.D{
				primitive.E{
					Key: "$inc",
					Value: bson.D{
						primitive.E{
							Key:   StockAmount,
							Value: 0 - amounts[id],
						},
					},
				},
			}

			_, err := o.StockCollection.UpdateOne(sessCtx, query, add)
			if err != nil {
				return nil, err
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
