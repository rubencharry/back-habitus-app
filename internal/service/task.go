package service

import (
	"backend-habitus-app/internal/domain/task"
)

type TaskService struct {
	rp task.TaskRepository
}

func NewTaskService(rp task.TaskRepository) *TaskService {
	return &TaskService{
		rp: rp,
	}
}