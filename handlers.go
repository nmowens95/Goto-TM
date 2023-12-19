package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Comment     string `json:"comment"`
	Status      string `json:"status"`
}

var mu sync.Mutex

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// insert task into db
	insert := "INSERT INTO tasks (Name, Description, Status) VALUES (?, ?, ?)"
	_, err := DB.Exec(insert, task.Name, task.Description, task.Comment, task.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {

}

// get task by id
func GetTask(w http.ResponseWriter, r *http.Request) {

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

}
