package service

import habitlog "backend-habitus-app/internal/domain/habit_log"

type HabitLogService struct {
	rp habitlog.HabitLogRepository
}

func NewHabitLogService(rp habitlog.HabitLogRepository) *HabitLogService {
	return &HabitLogService{
		rp: rp,
	}
}