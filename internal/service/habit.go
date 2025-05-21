package service

import (
	"backend-habitus-app/internal/domain/habit"
	"backend-habitus-app/internal/model"
)

type HabitService struct {
	rp habit.HabitRepository
}

func NewHabitService(rp habit.HabitRepository) *HabitService {
	return &HabitService{
		rp: rp,
	}
}

func (s *HabitService) GetAll() ([]*model.HabitDoc, error) {
	
	return []*model.HabitDoc{}, nil
}

func (s *HabitService) GetByID(id int) (*model.HabitDoc, error) {
	
	return &model.HabitDoc{}, nil
}

func (s *HabitService) Create(habit *model.HabitDoc) (*model.HabitDoc, error) {
	
	return &model.HabitDoc{}, nil
}

func (s *HabitService) Update(id int, habit *model.HabitDoc) (*model.HabitDoc, error) {
	
	return &model.HabitDoc{}, nil
}

func (s *HabitService) Delete(id int) error {
	return nil
}