package handler

import (
	"backend-habitus-app/internal/domain/habit"
	"net/http"
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


}

func (h *HabitHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *HabitHandler) GetByID(w http.ResponseWriter, r *http.Request) {

}

func (h *HabitHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *HabitHandler) Delete(w http.ResponseWriter, r *http.Request) {

}