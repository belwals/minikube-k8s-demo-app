package repository

import (
	"context"
	"errors"

	"github.com/belwals/minikube-k8s-demo-app/service/tiny-url-implementation/config"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client config.MongoClient

type ITinyUrlRepository interface {
	GetFullUrl(ctx context.Context, dbName string, tinyIdentifier string) (string, error)
	GenerateShortUrl(ctx context.Context, dbName string, fullUrl string) (string, error)
	IsShortUrlAlreadyGenerated(ctx context.Context, dbName string, fullUrl string) (string, error)
}

func (repo Client) GetFullUrl(ctx context.Context, dbName string, tinyIdentifier string) (string, error) {

	collection := repo.Client.Database(dbName).Collection("TinyUrl")

	// Find and remove the oldest document (FIFO)
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.M{"key": tinyIdentifier}).Decode(&result)
	if err != nil {
		return "", err
	}
	record := result["url"]
	return record.(string), nil
}

func (repo Client) GenerateShortUrl(ctx context.Context, dbName string, fullUrl string) (string, error) {
	collection := repo.Client.Database(dbName).Collection("TinyUrl")

	// Find and remove the oldest document (FIFO)
	newKey := uuid.New()
	input := bson.M{
		"url": fullUrl,
		"key": newKey.String(),
	}
	result, err := collection.InsertOne(context.TODO(), input)
	if err != nil || result.InsertedID == nil {
		return "", err
	}

	return newKey.String(), nil
}

func (repo Client) IsShortUrlAlreadyGenerated(ctx context.Context, dbName string, fullUrl string) (string, error) {
	collection := repo.Client.Database(dbName).Collection("TinyUrl")
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.M{"url": fullUrl}).Decode(&result)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return "", err
	}
	if errors.Is(err, mongo.ErrNoDocuments) || result == nil {
		return "", nil
	}

	record := result["key"]
	return record.(string), nil
}
