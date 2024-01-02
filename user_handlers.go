package main

import (
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
		http.Redirect(w, r, "/dashboard", http.StatusFound)
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

	http.Redirect(w, r, "/login", http.StatusFound)
}
