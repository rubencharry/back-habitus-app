package habitlog

import "backend-habitus-app/internal/model"

type HabitLogService interface {
	GetAll() ([]*model.HabitLogDoc, error)
}