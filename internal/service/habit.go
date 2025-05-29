package service

import (
	"backend-habitus-app/internal/domain/habit"
	"backend-habitus-app/internal/model"
	"errors"
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
	
	habits, err := s.rp.GetAll()
	if err != nil {
		return nil, err
	}

	var habitDocs []*model.HabitDoc
	for _, h := range habits {
		habitDoc := model.HabitToDoc(h)
		habitDocs = append(habitDocs, habitDoc)
	}

	return habitDocs, nil
}

func (s *HabitService) GetByID(id int) (*model.HabitDoc, error) {
	
	habit, err := s.rp.GetByID(id)
	if err != nil {
		return nil, err
	}

	return model.HabitToDoc(habit), nil
}

func (s *HabitService) Create(habit *model.HabitDoc) (*model.HabitDoc, error) {
	
	habitToModel := model.DocToHabit(habit)
	createdHabit, err := s.rp.Create(habitToModel)
	if err != nil {
		return nil, err
	}
	return model.HabitToDoc(createdHabit), nil
}

func (s *HabitService) Update(id int, habit *model.HabitDoc) (*model.HabitDoc, error) {
	if id < 0 {
		return nil, errors.New("invalid habit ID")
	}

	habitModel, err := s.rp.GetByID(id)
	if err != nil {
		return nil, err
	}

	habitModel.UpdateModel(habit)
	habitModel.ID = id

	_, err = s.rp.Update(id, habitModel)
	if err != nil {
		return nil, err
	}

	*habit = *model.HabitToDoc(habitModel)

	return habit, nil
}

func (s *HabitService) Delete(id int) error {

	if err := s.rp.Delete(id); err != nil {
		return err
	}

	return nil
}