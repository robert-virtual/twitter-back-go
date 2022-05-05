package main

import "fmt"

func findUser(id string) (*user, error) {
	var user user
	rows := db.QueryRow("SELECT * FROM users where id = ?", id)

	if err := rows.Scan(&user.Id, &user.UserName, &user.Name, &user.Password); err != nil {
		return nil, fmt.Errorf("images %q:%v", id, err)
	}
	return &user, nil

}
