package model

type TinyUrlResponse struct {
	Url         string `json:"url,omitempty"`
	ShortUrlKey string `json:"shortendUrl,omitempty"`
}
