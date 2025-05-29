package task

import "backend-habitus-app/internal/model"

type TaskService interface {
	GetAll() ([]*model.TaskDoc, error)
	GetByID(id int) (*model.TaskDoc, error)
	Create(habit *model.TaskDoc) (*model.TaskDoc, error)
	Update(id int, habit *model.TaskDoc) (*model.TaskDoc, error)
	Delete(id int) error
}