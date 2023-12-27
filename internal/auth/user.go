package auth

import (
	"github.com/nmowens95/Goto-TM/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserWithPassword(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// store hashed password in db
	_, err = database.DB.Exec("INSERT INTO users (Email, Password) Values (?, ?)", email, hashedPassword)
	return err
}

func AuthenticateUser(username, email, password string) (bool, error) {
	var hashedPassword string

	err := database.DB.QueryRow("SELECT Password From users WHERE Email = ?", email).Scan(&hashedPassword)
	if err != nil {
		return false, err // user not found
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err // passwords don't match
	}

	return true, nil // passwords match
}
