package handler

import "backend-habitus-app/internal/domain/task"

type TaskHandler struct {
	sv task.TaskService
}

func NewTaskHandler(sv task.TaskService) *TaskHandler {
	return &TaskHandler{
		sv: sv,
	}
}