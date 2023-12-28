package auth

import (
	"net/http"
)

func handlerUserLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	authenticated, err := AuthenticateUser(email, password)
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
