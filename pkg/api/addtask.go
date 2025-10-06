package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go1fl-final/pkg/db"
)

const dateLayout = "20060102"

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
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

	if task.Title == "" {
		writeError(w, "Task title is required")
		return
	}

	if err := checkDate(&task); err != nil {
		writeError(w, err.Error())
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeError(w, fmt.Sprintf("Failed to add task: %v", err))
		return
	}

	writeJSON(w, map[string]string{"id": fmt.Sprintf("%d", id)})
}

func checkDate(task *db.Task) error {
	now := time.Now()
	nowStr := now.Format(dateLayout)

	if task.Date == "" {
		task.Date = nowStr
		return nil
	}

	t, err := time.Parse(dateLayout, task.Date)
	if err != nil {
		return fmt.Errorf("invalid date format")
	}

	nowDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	tDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)

	if task.Repeat != "" {
		_, err := db.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return fmt.Errorf("invalid repeat rule: %v", err)
		}

		if tDay.Before(nowDay) {
			next, _ := db.NextDate(now, task.Date, task.Repeat)
			task.Date = next
		}
	} else {
		if tDay.Before(nowDay) {
			task.Date = nowStr
		}
	}

	return nil
}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
