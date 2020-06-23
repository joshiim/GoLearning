package db

import (
	"/Users/Santosh Joshi/Documents/goWorkspace/src/github.com/mongodb/mongo-go-driver/mongo"
	"context"
)

func GetDBCollection() (*mongo.Collection, error) {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	collection := client.Database("GoLogin").Collection("users")
	return collection, nil
}