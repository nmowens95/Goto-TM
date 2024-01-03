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
	router.Get("/tasks", handlerGetTasks)
	router.Get("/tasks/{id}", GetTask)
	router.Post("/tasks", handlerCreateTask)
	router.Put("/tasks/{id}", handlerUpdateTask)
	router.Delete("/tasks/{id}", handlerDeleteTask)

	// Users
	apiRouter := chi.NewRouter()
	apiRouter.Put("/signup", handlerUserSignup)
	apiRouter.Post("/users", handlerCreateUser)
	apiRouter.Post("/login", handlerUserLogin)
	router.Mount("/api", apiRouter)

	srv := &http.Server{
		Addr:    ":" + "8080",
		Handler: router,
	}

	defer database.DB.Close()

	log.Printf("Listening on port :%v", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
