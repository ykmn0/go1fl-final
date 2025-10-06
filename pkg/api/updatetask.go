package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go1fl-final/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, "Failed to read request body")
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &task); err != nil {
		writeError(w, "Invalid JSON format")
		return
	}

	if task.ID == "" {
		writeError(w, "Task ID is required")
		return
	}

	if task.Title == "" {
		writeError(w, "Task title is required")
		return
	}

	if err := checkDate(&task); err != nil {
		writeError(w, err.Error())
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		writeError(w, fmt.Sprintf("Failed to update task: %v", err))
		return
	}

	writeJSON(w, map[string]interface{}{})
}
