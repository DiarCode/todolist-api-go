package main

//https://github.com/NikSchaefer/go-fiber

import (
	"log"

	"github.com/DiarCode/todo-go-api/pkg/server"
)

const PORT = "8080"

func main() {

	handlers := new(server.Handler)
	routes := handlers.InitRoutes()
	srv := new(server.Server)

	log.Printf("Server started on port: %s\n", PORT)

	if err := srv.Run(PORT, routes); err != nil {
		log.Fatalf("Error occured while running server: %s\n", err.Error())
	}

}
