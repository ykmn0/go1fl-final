package api

import (
	"fmt"
	"net/http"
	"time"

	"go1fl-final/pkg/db"
)

func taskDoneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		writeError(w, "Task ID is required")
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeError(w, err.Error())
		return
	}

	if task.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			writeError(w, fmt.Sprintf("Failed to delete task: %v", err))
			return
		}
	} else {
		now := time.Now()
		next, err := db.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeError(w, fmt.Sprintf("Failed to calculate next date: %v", err))
			return
		}

		if err := db.UpdateTaskDate(id, next); err != nil {
			writeError(w, fmt.Sprintf("Failed to update task date: %v", err))
			return
		}
	}

	writeJSON(w, map[string]interface{}{})
}
