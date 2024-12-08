package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnectInput struct {
	userName   string
	password   string
	clusterUrl string
}

func NewMongoInput(username, password, clusterUrl string) MongoConnectInput {
	return MongoConnectInput{
		userName: username, password: password, clusterUrl: clusterUrl,
	}
}

type MongoClient struct {
	Client *mongo.Client
}

// func (c MongoClient) GetClient() *mongo.Client {
// 	if c.client == nil {
// 		panic("mongo client is not initialized")
// 	}
// 	return c.client
// }

func (input MongoConnectInput) NewClient(ctx context.Context) (*MongoClient, error) {

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s",
		input.userName, input.password, input.clusterUrl,
	)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	return &MongoClient{
		Client: client,
	}, nil
}
