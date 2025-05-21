package main

import (
	"backend-habitus-app/cmd/app"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg := &app.ConfigServerMySQLChi{
		Db: &mysql.Config{
			User:   "root",
			Passwd: os.Getenv("MYSQL_PASSWORD"),
			Net:    "tcp",
			Addr:   "localhost:3306",
			DBName: os.Getenv("MYSQL_DATABASE_NAME"),
		},
		ServerAddress: "127.0.0.1:8080",
	}

	appServer := app.NewConfigServerMySQLChi(cfg)

	if err := appServer.Setup(); err != nil {
		log.Fatalf("Error setting up server: %v", err)
	}

	if err := appServer.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	fmt.Print("Server started on 8080", "\n")
}
