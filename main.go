package main

import (
	"github.com/gpark1005/MovieProject/handlers"
	"github.com/gpark1005/MovieProject/repository"
	"github.com/gpark1005/MovieProject/service"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	fn := "/home/gp/GolandProjects/MovieProject/moviedb.json"

	ext := filepath.Ext(fn)

	if ext != ".json" {
		log.Fatalln("File extension invalid")
	}

	r := repository.NewRepository(fn)

	svc := service.NewService(r)

	hdlr := handlers.NewMovieHandler(svc)

	router := handlers.ConfigureRouter(hdlr)

	svr := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	log.Fatalln(svr.ListenAndServe())

}
