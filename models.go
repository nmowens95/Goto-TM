package main

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Comment     string `json:"comment"`
	Status      string `json:"status"`
	UserID      int    `json:"userid"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
