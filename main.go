package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	createConnection()
	router := gin.Default()
	posts := router.Group("/posts")
	{
		posts.GET("", getPosts)

	}
	port := os.Getenv("PORT")
	router.Run(":" + port)
}
