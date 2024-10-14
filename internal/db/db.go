package db

import (
	"database/sql"
	"log"
)

func InitRDB() *sql.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/invoice?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("could not ping the database: %v", err)
	}

	log.Println("Connected to the database")
	return db
}
