package service

import (
	"context"

	"github.com/belwals/minikube-k8s-demo-app/service/tiny-url-implementation/model"
	"github.com/belwals/minikube-k8s-demo-app/service/tiny-url-implementation/repository"
)

const TinyUrlDatabaseName = "tiny-url"

type IRestService interface {
	GetFullURl(ctx context.Context, uniqueId string) (string, error)
	GenerateShortUrl(ctx context.Context, url string) (string, error)
}

type TinuUrlService struct {
	env  model.Environment
	repo repository.ITinyUrlRepository
}

func NewTinyUrlService(env model.Environment, repo repository.ITinyUrlRepository) TinuUrlService {
	return TinuUrlService{
		env:  env,
		repo: repo,
	}
}

func (service TinuUrlService) GetFullURl(ctx context.Context, uniqueId string) (string, error) {
	return service.repo.GetFullUrl(ctx, TinyUrlDatabaseName, uniqueId)
}
func (service TinuUrlService) GenerateShortUrl(ctx context.Context, url string) (string, error) {
	shortKey, err := service.repo.IsShortUrlAlreadyGenerated(ctx, TinyUrlDatabaseName, url)
	if err != nil {
		return "", err
	}
	if len(shortKey) != 0 {
		return shortKey, nil
	}
	// we didn't have record already in system hence creating a new one
	return service.repo.GenerateShortUrl(ctx, TinyUrlDatabaseName, url)
}
