package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func createConnection() {
	// cfg == config
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "ec2-3-141-200-64.us-east-2.compute.amazonaws.com",
		DBName: "twittergo",
	}
	var error error
	db, error = sql.Open("mysql", cfg.FormatDSN())
	if error != nil {
		log.Fatal(error)
	}
	pingError := db.Ping()
	if pingError != nil {
		log.Fatal(pingError)
	}
	fmt.Println("Conectado a base de datos")
}
