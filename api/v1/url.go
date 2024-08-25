package v1

import (
	"fmt"
	"net/http"
	"panvalkar1994/furle/models"
	"panvalkar1994/furle/services"

	"github.com/gin-gonic/gin"
)

func Shorten(c *gin.Context) {
	var url models.ShortenRequest

	if err := c.BindJSON(&url); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Printf("URL: %s\n", url.Url)

	shortened, err := services.SaveShortenUrl(url.Url)
	if err != nil {
		c.JSON(417, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shortened": shortened,
	})

}
