package api

import (
	api_v1 "panvalkar1994/furle/api/v1"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	api_v1.Routes(v1)
}
