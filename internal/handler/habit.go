package handler

import (
	"backend-habitus-app/internal/domain/habit"
	"backend-habitus-app/internal/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type HabitHandler struct {
	sv habit.HabitService
}

func NewHabitHandler(sv habit.HabitService) *HabitHandler {
	return &HabitHandler{
		sv: sv,
	}
}

func (h *HabitHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	habits, err := h.sv.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(habits); err != nil {
		http.Error(w, "Failed to encode habits", http.StatusInternalServerError)
		return
	}
}

func (h *HabitHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid habit ID", http.StatusBadRequest)
		return
	}

	habit, err := h.sv.GetByID(id)
	if err != nil {
		if err.Error() == "habit not found" {
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

func (h *HabitHandler) Create(w http.ResponseWriter, r *http.Request) {
	var habit model.HabitDoc
	if err := json.NewDecoder(r.Body).Decode(&habit); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createdHabit, err := h.sv.Create(&habit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdHabit); err != nil {
		http.Error(w, "Failed to encode created habit", http.StatusInternalServerError)
		return
	}
}

func (h *HabitHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid habit ID", http.StatusBadRequest)
		return
	}

	var habit model.HabitDoc
	if err := json.NewDecoder(r.Body).Decode(&habit); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	updatedHabit, err := h.sv.Update(id, &habit)

	if err != nil {
		if err.Error() == "habit not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(updatedHabit); err != nil {
		http.Error(w, "Failed to encode updated habit", http.StatusInternalServerError)
		return
	}

}

func (h *HabitHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid habit ID", http.StatusBadRequest)
		return
	}

	err = h.sv.Delete(id)
	if err != nil {
		if err.Error() == "habit not found" {
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
