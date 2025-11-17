package main

import (
	"database/sql"
	"gin-rest-api/internal/database"
	"gin-rest-api/internal/env"
	"log"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	port      int
	models    database.Models
	jwtSecret string
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)

	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		models:    models,
		jwtSecret: env.GetEnvString("JWT_SECRET", "defaultsecret"),
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
