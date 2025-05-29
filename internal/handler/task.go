package handler

import (
	"backend-habitus-app/internal/domain/task"
	"backend-habitus-app/internal/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	sv task.TaskService
}

func NewTaskHandler(sv task.TaskService) *TaskHandler {
	return &TaskHandler{
		sv: sv,
	}
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	tasks, err := h.sv.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Failed to encode tasks", http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid habit ID", http.StatusBadRequest)
		return
	}

	habit, err := h.sv.GetByID(id)
	if err != nil {
		if err.Error() == "task not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(habit); err != nil {
		http.Error(w, "Failed to encode habit", http.StatusInternalServerError)
		return
	}

}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var task model.TaskDoc
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createdTask, err := h.sv.Create(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdTask); err != nil {
		http.Error(w, "Failed to encode created task", http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task model.TaskDoc
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	updatedTask, err := h.sv.Update(id, &task)

	if err != nil {
		if err.Error() == "task not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(updatedTask); err != nil {
		http.Error(w, "Failed to encode updated task", http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = h.sv.Delete(id)
	if err != nil {
		if err.Error() == "task not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(nil); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}