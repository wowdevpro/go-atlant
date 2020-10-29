package main

import (
	"context"
	"github.com/wowdevpro/go-atlant/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

var collection *mongo.Collection
var ctx = context.TODO()

func saveCsvData(data [][]string) error {
	client, err := getClient()
	if err != nil {
		return err
	}

	collection = client.Database("csv_service").Collection("products")
	var price int

	for _, item := range data {
		price, _ = strconv.Atoi(item[1])

		filter := bson.M{"Name": item[0]}

		update := bson.M{
			"$set": bson.M{
				"LastPrice": price,
				"UpdatedAt": time.Now().Unix(),
			},
			"$inc": bson.M{
				"CountUpdates": 1,
			},
		}

		upsert := true
		opt := options.FindOneAndUpdateOptions{
			Upsert: &upsert,
		}

		collection.FindOneAndUpdate(ctx, filter, update, &opt)
	}

	client.Disconnect(ctx)

	return nil
}

func getProducts(orderBy string, size int32, number int32) (products []*proto.Product, err error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	collection = client.Database("csv_service").Collection("products")

	skip := int64(size * (number - 1))
	limit := int64(size)

	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
		Sort:  bson.D{{orderBy, -1}},
	}

	cur, err := collection.Find(ctx, bson.M{}, &opts)
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &products)
	if err != nil {
		return nil, err
	}

	err = cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	client.Disconnect(ctx)

	return products, nil
}

func getClient() (*mongo.Client, error) {
	uri := "mongodb://root:secret@mongo:27017/"

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
