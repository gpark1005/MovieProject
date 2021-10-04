package main

import (
	"MovieProject/handlers"
	"MovieProject/repo"
	"MovieProject/service"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	file := "moviedb.json"

	ext := filepath.Ext(file)
	if ext != ".json" {
		log.Fatal("Invalid File Extension")
	}

	repository := repo.NewRepo(file)
	srv := service.ForService(repository)
	mvHandler := handlers.NewMovieHandler(srv)
	router := handlers.ConfigureRouter

	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	log.Fatal(server.ListenAndServe())
}
