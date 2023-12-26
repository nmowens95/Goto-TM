package main

import (
	"encoding/json"
	"net/http"

	"github.com/nmowens95/Goto-TM/internal/database"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addUser := "INSERT INTO user (Email) VALUES (?)"
	result, err := database.DB.Exec(addUser, user.Email)
	if err != nil {
		http.Error(w, "There was an issue adding this user", http.StatusInternalServerError)
		return
	}

	// Get ID of the las inserted user
	userID, _ := result.LastInsertId()
	user.ID = int(userID) // Assign inserted user's Id to the user Struct

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
