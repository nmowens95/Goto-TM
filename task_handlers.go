package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"
)

var mu sync.Mutex

// Create a single task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// insert task into db
	insert := "INSERT INTO tasks (Name, Description, Comment, Status, UserID) VALUES (?, ?, ?, ?, ?)"
	_, err := DB.Exec(insert, task.Name, task.Description, task.Comment, task.Status)
	if err != nil {
		http.Error(w, "Not able to insert task, something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// get individual task (by ID)
func GetTask(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	mu.Lock()
	defer mu.Unlock()

	var task Task
	err = DB.QueryRow("SELECT ID, Name, Description, Comment, Status FROM tasks WHERE ID = ?", id).Scan(&task.ID, &task.Name, &task.Description, &task.Comment, &task.Status, &task.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	rows, err := DB.Query("SELECT ID, Name, Description, Comment, Status, UserID FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.Comment, &task.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Update a current task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		http.Error(w, "Task not found", http.StatusBadRequest)
		return
	}

	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	result, err := DB.Exec("UPDATE tasks SET Name = ?, Description = ?, Comment = ?, Status = ? WHERE ID = ?", updatedTask.Name, updatedTask.Description, updatedTask.Comment, updatedTask.Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAff, _ := result.RowsAffected()
	if rowsAff == 0 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete a created task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		http.Error(w, "Task ID not found", http.StatusNotFound)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	result, err := DB.Exec("DELETE FROM tasks WHERE ID = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAff, _ := result.RowsAffected()
	if rowsAff == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
