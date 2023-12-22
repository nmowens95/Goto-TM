package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	openDB()
	if DB == nil {
		log.Fatal("DB is not initialized properly")
	}

	port := "8080"
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Handle("/", http.FileServer(http.Dir(".")))
	router.Get("/tasks", GetTasks)
	router.Get("/tasks/{id}", GetTask)
	router.Post("/tasks", CreateTask)
	router.Put("/tasks/{id}", UpdateTask)
	router.Delete("/tasks/{id}", DeleteTask)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	defer DB.Close()

	log.Print("Listening...")
	log.Fatal(srv.ListenAndServe())
}
