package routes

import (
	"panvalkar1994/furle/api"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	api_group := r.Group("/api")
	api.Routes(api_group)
}
