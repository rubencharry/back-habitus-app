package handler

import (
	"backend-habitus-app/internal/domain/habit_log"
	"backend-habitus-app/internal/model"
	"encoding/json"
	"net/http"
)

type HabitLogHandler struct {
	sv habitlog.HabitLogService
}

func NewHabitLogHandler(sv habitlog.HabitLogService) *HabitLogHandler {
	return &HabitLogHandler{
		sv: sv,
	}
}

func (h *HabitLogHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	habitLogs, err := h.sv.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(habitLogs); err != nil {
		http.Error(w, "Failed to encode habit logs", http.StatusInternalServerError)
		return
	}
}

func (h *HabitLogHandler) Create(w http.ResponseWriter, r *http.Request) {
	var habitLog model.HabitLogDoc
	if err := json.NewDecoder(r.Body).Decode(&habitLog); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createdHabit, err := h.sv.Create(&habitLog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdHabit); err != nil {
		http.Error(w, "Failed to encode created habit log", http.StatusInternalServerError)
		return
	}
}