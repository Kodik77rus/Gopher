package main

import (
	"github.com/Kodik77rus/todo-app"
	"github.com/Kodik77rus/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running server: %s", err.Error())
	}
}
