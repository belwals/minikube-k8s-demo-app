package repository

import (
	"github.com/belwals/minikube-k8s-demo-app/service/tiny-url-implementation/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ITinyUrlRepository interface {
	GetFullUrl(tinyIdentifier string) (string, error)
	GenerateShortUrl(fullUrl string) (string, error)
}

type TinyUrlRepository struct {
	client *mongo.Client
}

func NewTinyUrlRepo(env model.Environment) (*TinyUrlRepository, error) {
	// client, err := mongo.Connect(context.TODO(), options.Client().
	// 	ApplyURI(env.MongoDbUri))
	// if err != nil {
	// 	return nil, err
	// }

	// return &TinyUrlRepository{
	// 	client: client,
	// }, nil
	return &TinyUrlRepository{}, nil
}

func (repo TinyUrlRepository) GetFullUrl(tinyIdentifier string) (string, error) {
	// databaseName := ""
	// repo.client.Database(databaseName).Collection(tinyIdentifier)
	return "NOT IMPEMENTED", nil
}

func (repo TinyUrlRepository) GenerateShortUrl(fullUrl string) (string, error) {
	// databaseName := ""
	// repo.client.Database(databaseName).Collection(tinyIdentifier)
	return "NOT IMPEMENTED", nil
}
