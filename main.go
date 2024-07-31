package main

import (
	"database/sql"
	"infilion-assignment/internal"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	Db     *sql.DB
	Engine *gin.Engine
	err    error
)

func main() {
	Engine = gin.Default()
	//enable CORS
	Engine.Use(cors.New(internal.CORS()))
	//DB client
	Db, err = internal.DBConnection()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Connected to database successfully")
	
	app := internal.MyApp{
		DB: Db,
	}
	app.Routes(Engine)
	Engine.Run()
}
