package main

import (
	"animeList/internal/content"
	"animeList/internal/content/postgres"
	"animeList/internal/http"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	url := os.Getenv("URL")

	db := postgres.NewDB()
	if err := db.Connect(url); err != nil {
		panic(err)
	}
	defer func(database content.Database) {
		err := database.Close()
		if err != nil {

		}
	}(db)

	//Creating and run new server
	srv := http.NewServer(context.Background(), ":8080", db)
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

	srv.WaitForGT()
}
