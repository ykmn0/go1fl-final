package api

import (
	"net/http"

	"go1fl-final/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")

	var tasks []*db.Task
	var err error

	if search != "" {
		tasks, err = db.SearchTasks(search, 50)
	} else {
		tasks, err = db.Tasks(50)
	}

	if err != nil {
		writeError(w, "Failed to fetch tasks")
		return
	}

	if tasks == nil {
		tasks = make([]*db.Task, 0)
	}

	writeJSON(w, TasksResp{Tasks: tasks})
}
