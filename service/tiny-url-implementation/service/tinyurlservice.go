package service

import (
	"github.com/belwals/minikube-k8s-demo-app/service/tiny-url-implementation/model"
	"github.com/belwals/minikube-k8s-demo-app/service/tiny-url-implementation/repository"
)

type IRestService interface {
	GetFullURl(uniqueId string) (string, error)
	GenerateShortUrl(url string) (string, error)
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

func (service TinuUrlService) GetFullURl(uniqueId string) (string, error) {
	// TODO: dummy implementation
	return "Dummy URL", nil
}
func (service TinuUrlService) GenerateShortUrl(url string) (string, error) {
	// TODO: dummy implementation
	return "Dummy Generated", nil
}
