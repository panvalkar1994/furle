package main

import (
	"panvalkar1994/furle/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "panvalkar1994/furle/db"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	r := gin.Default()

	r.Static("/public", "./public")
	r.LoadHTMLGlob("public/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	routes.Routes(r)
	r.Run()
}
