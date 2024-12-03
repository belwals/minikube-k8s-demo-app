package service

type IRestService interface {
	GetFullURl(uniqueId string) (string, error)
	GenerateShortUrl(url string) (string, error)
}
