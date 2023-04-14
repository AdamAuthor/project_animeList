package main

import (
	"animeList/internal/content/inmemory"
	"animeList/internal/http"
	"context"
	"fmt"
)

func main() {
	store := inmemory.NewDB()

	srv := http.NewServer(context.Background(), ":8080", store)
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

	srv.WaitForGT()
}
