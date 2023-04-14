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

	store := postgres.NewDB()
	if err := store.Connect(url); err != nil {
		panic(err)
	}
	defer func(store content.Content) {
		err := store.Close()
		if err != nil {
			
		}
	}(store)

	//Creating and run new server
	srv := http.NewServer(context.Background(), ":8080", store)
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

	srv.WaitForGT()
}
