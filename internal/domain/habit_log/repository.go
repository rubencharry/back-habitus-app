package habitlog

import "backend-habitus-app/internal/model"

type HabitLogRepository interface {
	GetAll() ([]*model.HabitLog, error)
}