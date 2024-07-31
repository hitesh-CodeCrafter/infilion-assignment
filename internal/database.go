package internal

import (
	"database/sql"
	"log"
)

func DBConnection() (db *sql.DB, err error) {
	dsn := "username:password@tcp(127.0.0.1:3306)/dbname" // database url
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return db, err
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return db, err
	}
	return db, nil
}
