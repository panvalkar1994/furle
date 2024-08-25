package models

type ShortenRequest struct {
	Url string `json:"url"`
}

type ShortUrl struct {
	Url      string `json:"url"`
	ShortUrl string `json:"short_url"`
}
