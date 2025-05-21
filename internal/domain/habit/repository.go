package habit

import "backend-habitus-app/internal/model"

type HabitRepository interface {
	GetAll() ([]*model.Habit, error)
	GetByID(id int) (*model.Habit, error)
	Create(habit *model.Habit) (*model.Habit, error)
	Update(id int,habit *model.Habit) (*model.Habit, error)
	Delete(id int) error
}