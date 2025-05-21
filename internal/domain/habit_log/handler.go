package habitlog

import "net/http"

type HabitLogHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}