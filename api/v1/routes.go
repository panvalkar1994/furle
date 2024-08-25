package v1

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.GET("/ping", Ping)
	r.POST("/shorten", Shorten)
}
