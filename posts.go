package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//ctx == context
func getPosts(ctx *gin.Context) {
	var posts []post
	// solicitar datos a base de datos
	posts, err := findPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	//enviar datos a cliente
	ctx.JSON(http.StatusOK, posts) //enviar datos a cliente
}

func findPosts() ([]post, error) {
	var posts []post
	rows, error := db.Query("SELECT * FROM posts")
	if error != nil {
		return nil, fmt.Errorf("posts %v", error)
	}
	defer rows.Close()
	for rows.Next() {
		var post post
		if err := rows.Scan(&post.Id, &post.Content, &post.UserId, &post.CreatedAt); err != nil {
			return nil, fmt.Errorf("posts %v", err)
		}
		if images, err := findImages(post.Id); err == nil {
			post.Images = &images
		}
		if user, err := findUser(post.UserId); err == nil {
			post.User = user
		}
		posts = append(posts, post)

	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("posts %v", err)
	}
	return posts, nil

}
