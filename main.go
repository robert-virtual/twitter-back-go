package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

var argon = argon2.DefaultConfig()

func main() {
	createConnection()
	router := gin.Default()
	posts := router.Group("/posts")
	{
		posts.GET("", getPosts)
		posts.POST("", postPost)

	}
	users := router.Group("/users")
	{
		users.GET("", getUsers)
		users.POST("", postUser)

	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
