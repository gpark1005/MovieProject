package main

import (
	"MovieProject/handlers"
	"log"
)

func main() {
	server := handlers.NewServer()

	log.Fatal(server.ListenAndServe())
}
