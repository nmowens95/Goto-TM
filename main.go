package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	openDB()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Handle("/", http.FileServer(http.Dir(".")))
	router.Get("/tasks", GetTasks)
	router.Get("/tasks/{id}", GetTask)
	router.Post("/tasks", CreateTask)
	router.Put("/tasks/{id}", UpdateTask)
	router.Delete("/tasks/{id}", DeleteTask)

	apiRouter := chi.NewRouter()
	router.Mount("/api", apiRouter)

	srv := &http.Server{
		Addr:    ":" + "8080",
		Handler: router,
	}

	defer DB.Close()

	log.Print("Listening...")
	log.Fatal(srv.ListenAndServe())
}
