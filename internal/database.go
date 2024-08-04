package internal

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (db *sql.DB, err error) {
	dsn := "root:Hitesh0508@@tcp(127.0.0.1:3306)/cetec" // database url
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
