package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := "8080"
	router := chi.NewRouter()
	router.Handle("/", http.FileServer(http.Dir(".")))
	corsRouter := middlewareCors(router)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsRouter,
	}

	log.Print("Listening...")
	log.Fatal(srv.ListenAndServe())
}
