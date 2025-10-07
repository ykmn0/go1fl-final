package api

import (
	"fmt"
	"net/http"

	"go1fl-final/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeError(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	if err := db.DeleteTask(id); err != nil {
		writeError(w, fmt.Sprintf("Failed to delete task: %v", err), http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]interface{}{})
}
