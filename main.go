package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tahsin005/layered-based-architecture/todo-app/config"
	"github.com/tahsin005/layered-based-architecture/todo-app/database"
	"github.com/tahsin005/layered-based-architecture/todo-app/handler"
	"github.com/tahsin005/layered-based-architecture/todo-app/repository"
	"github.com/tahsin005/layered-based-architecture/todo-app/service"
)

func main() {
	cfg := config.LoadConfig()
	db := database.InitDB(cfg)
	defer db.Close()


	repo := repository.NewTodoRepository(db)
	svc := service.NewTodoService(repo)


	r := mux.NewRouter()
	handler.NewTodoHandler(r, svc)

	log.Println("Starting server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}