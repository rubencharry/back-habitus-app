package habit

import "backend-habitus-app/internal/model"

type HabitService interface {
	GetAll() ([]*model.HabitDoc, error)
	GetByID(id int) (*model.HabitDoc, error)
	Create(habit *model.HabitDoc) (*model.HabitDoc, error)
	Update(id int,habit *model.HabitDoc) (*model.HabitDoc, error)
	Delete(id int) error
}