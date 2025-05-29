package habitlog

import "backend-habitus-app/internal/model"

type HabitLogRepository interface {
	GetAll() ([]*model.HabitLog, error)
	GetByID(id int) (*model.HabitLog, error)
	Create(habitLog *model.HabitLog) (*model.HabitLog, error)
	Update(id int, habitLog *model.HabitLog) (*model.HabitLog, error)
}

