package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nmowens95/Goto-TM/internal/database"
)

func main() {
	database.OpenDB()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Handle("/", http.FileServer(http.Dir(".")))

	// Tasks
	router.Get("/tasks", GetTasks)
	router.Get("/tasks/{id}", GetTask)
	router.Post("/tasks", CreateTask)
	router.Put("/tasks/{id}", UpdateTask)
	router.Delete("/tasks/{id}", DeleteTask)

	// Users
	apiRouter := chi.NewRouter()
	apiRouter.Post("/users", CreateUser)
	router.Mount("/api", apiRouter)

	srv := &http.Server{
		Addr:    ":" + "8080",
		Handler: router,
	}

	defer database.DB.Close()

	log.Print("Listening...")
	log.Fatal(srv.ListenAndServe())
}
