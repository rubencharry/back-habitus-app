package task

import "backend-habitus-app/internal/model"

type TaskRepository interface {
	GetAll() ([]*model.Task, error)
	GetByID(id int) (*model.Task, error)
	Create(habit *model.Task) (*model.Task, error)
	Update(id int, habit *model.Task) (*model.Task, error)
	Delete(id int) error
}