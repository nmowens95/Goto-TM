package main

type DBStructure struct {
	Task  map[int]Task
	Users map[int]Users
}

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Comment     string `json:"comment"`
	Status      string `json:"status"`
}

type Users struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}
