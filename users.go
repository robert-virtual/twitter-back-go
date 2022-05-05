package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(ctx *gin.Context) {
	var users []user
	// solicitar datos a base de datos
	users, err := findUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	//enviar datos a cliente
	ctx.JSON(http.StatusOK, users) //enviar datos a cliente
}

func postUser(ctx *gin.Context) {
	var user user

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	rows, err := createUser(user)
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

func findUsers() ([]user, error) {
	var users []user
	rows, error := db.Query("SELECT * FROM users ORDER BY createdAt DESC")
	if error != nil {
		return nil, fmt.Errorf("find users %v", error)
	}
	defer rows.Close()
	for rows.Next() {
		var user user
		if err := rows.Scan(&user.Id, &user.Name, &user.UserName, &user.Password, &user.CreatedAt); err != nil {
			return nil, fmt.Errorf("posts %v", err)
		}
		user.Password = nil
		users = append(users, user)

	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("posts %v", err)
	}
	return users, nil

}

func createUser(user user) (int64, error) {
	password, err := argon.HashEncoded([]byte(*user.Password))
	if err != nil {
		return 0, fmt.Errorf("create post:%v", err)
	}
	res, err := db.Exec("insert into users(name,userName,password) values (?,?,?)", user.Name, user.UserName, password)
	if err != nil {
		return 0, fmt.Errorf("create post:%v", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("create post:%v", err)
	}

	return rows, err
}

func findUser(id string) (*user, error) {
	var user user
	rows := db.QueryRow("SELECT * FROM users where id = ?", id)

	if err := rows.Scan(&user.Id, &user.UserName, &user.Name, &user.Password, &user.CreatedAt); err != nil {
		return nil, fmt.Errorf("images %q:%v", id, err)
	}
	return &user, nil

}
