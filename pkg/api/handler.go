package api

import (
	"fmt"
	"net/http"
	"time"
)

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	repeat := r.FormValue("repeat")
	dstart := r.FormValue("date")
	if dstart == "" {
		http.Error(w, "parameter 'date' is required", http.StatusBadRequest)
		return
	}

	var now time.Time
	nowStr := r.FormValue("now")
	if nowStr == "" {
		now = time.Now()
	} else {
		var err error
		now, err = time.Parse(dateLayout, nowStr)
		if err != nil {
			http.Error(w, "invalid 'now' parameter", http.StatusBadRequest)
			return
		}
	}

	next, err := NextDate(now, dstart, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, next)
}
