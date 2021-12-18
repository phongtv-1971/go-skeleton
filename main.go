package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/phongtv-1971/go-skeleton/api"
	db "github.com/phongtv-1971/go-skeleton/db/sqlc"
	"github.com/phongtv-1971/go-skeleton/util"
	"io"
	"log"
	"os"
)

var environment = os.Getenv("APP_ENV")

func init() {
	if environment == "production" {
		gin.SetMode(gin.ReleaseMode)
		f, _ := os.Create("logs/production.log")
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	} else {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	config, err := util.LoadConfig(".", environment)
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
