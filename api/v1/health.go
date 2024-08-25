package v1

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.HTML(200, "ping.html", nil)
}
