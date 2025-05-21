package handler

import "backend-habitus-app/internal/domain/habit_log"

type HabitLogHandler struct {
	sv habitlog.HabitLogService
}

func NewHabitLogHandler(sv habitlog.HabitLogService) *HabitLogHandler {
	return &HabitLogHandler{
		sv: sv,
	}
}