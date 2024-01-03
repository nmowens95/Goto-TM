package main

import (
	"encoding/json"
	"net/http"

	"github.com/nmowens95/Goto-TM/internal/auth"
)

func handlerUserLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	authenticated, err := auth.AuthenticateUser(email, password)
	if err != nil {
		http.Error(w, "Error authenticating user", http.StatusInternalServerError)
		return
	}

	if authenticated {
		// Respond with success or redirect
		w.WriteHeader(http.StatusOK)
		response := map[string]string{"message": "User login successful!"}
		json.NewEncoder(w).Encode(response)
	} else {
		// User authentication failed
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	}
}

func handlerUserSignup(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	err := auth.CreateUserWithPassword(email, password)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "Signup Successful!"}
	json.NewEncoder(w).Encode(response)
}
