package service

import (
	"backend-habitus-app/internal/domain/task"
	"backend-habitus-app/internal/model"
	"errors"
)

type TaskService struct {
	rp task.TaskRepository
}

func NewTaskService(rp task.TaskRepository) *TaskService {
	return &TaskService{
		rp: rp,
	}
}

func (s *TaskService) GetAll() ([]*model.TaskDoc, error) {
	tasks, err := s.rp.GetAll()
	if err != nil {
		return nil, err
	}

	var taskDocs []*model.TaskDoc
	for _, t := range tasks {
		taskDoc := model.TaskToDoc(t)
		taskDocs = append(taskDocs, taskDoc)
	}

	return taskDocs, nil
}

func (s *TaskService) GetByID(id int) (*model.TaskDoc, error) {
	task, err := s.rp.GetByID(id)
	if err != nil {
		return nil, err
	}
	return model.TaskToDoc(task), nil
}

func (s *TaskService) Create(task *model.TaskDoc) (*model.TaskDoc, error) {
	taskToModel := model.DocToTask(task)
	createdTask, err := s.rp.Create(taskToModel)
	if err != nil {
		return nil, err
	}
	return model.TaskToDoc(createdTask), nil
}

func (s *TaskService) Update(id int, task *model.TaskDoc) (*model.TaskDoc, error) {
	if id < 0 {
		return nil, errors.New("invalid habit ID")
	}

	taskModel, err := s.rp.GetByID(id)
	if err != nil {
		return nil, err
	}

	taskModel.UpdateModel(task)
	taskModel.ID = id

	_, err = s.rp.Update(id, taskModel)
	if err != nil {
		return nil, err
	}

	*task = *model.TaskToDoc(taskModel)

	return task, nil
}

func (s *TaskService) Delete(id int) error {
	if err := s.rp.Delete(id); err != nil {
		return err
	}

	return nil
}
