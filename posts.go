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
			"error": err.Error(),
		})
		return
	}
	//enviar datos a cliente
	ctx.JSON(http.StatusOK, posts) //enviar datos a cliente
}
func postPost(ctx *gin.Context) {
	var post post

	if err := ctx.BindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	rows, err := createPost(post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"RowsAffected": rows,
	})
}

func findPosts() ([]post, error) {
	var posts []post
	rows, error := db.Query("SELECT * FROM posts ORDER BY createdAt DESC")
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
			user.Password = nil
			post.User = user
		}
		posts = append(posts, post)

	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("posts %v", err)
	}
	return posts, nil

}

func createPost(post post) (int64, error) {
	res, err := db.Exec("insert into posts(content,userId) values (?,?)", post.Content, post.UserId)
	if err != nil {
		return 0, fmt.Errorf("create post:%v", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("create post:%v", err)
	}

	return rows, err
}
