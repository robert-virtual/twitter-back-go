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
		posts.POST("", postPost)

	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
