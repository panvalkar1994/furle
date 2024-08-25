package services

import (
	"errors"
	"os"
	"panvalkar1994/furle/db"
	"panvalkar1994/furle/models"
	utils "panvalkar1994/furle/utils"
	"strings"
)

func SaveShortenUrl(url string) (string, error) {
	if url == "" {
		return "", errors.New("URL is required")
	}

	// Check if the url is already in the database
	if short, ok := db.Db[url]; ok {
		surl := (short).(models.ShortUrl)
		return surl.ShortUrl, nil
	}

	println("generating short url")
	shorturl := utils.SmallBatch.GetNextShortUrl()
	println("Short URL: ", shorturl)
	// TODO: Check if the short url is already in the database
	short := models.ShortUrl{
		ShortUrl: shorturl,
		Url:      url,
	}

	// Save the url in the database
	db.Db[url] = short

	shorturl = strings.ToLower(os.Getenv("site")) + "/" + shorturl

	return shorturl, nil
}
