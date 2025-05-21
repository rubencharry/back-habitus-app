package repository

import (
	"backend-habitus-app/internal/model"
	"database/sql"
)

type HabitRepository struct {
	db *sql.DB
}

func NewHabitRepository(db *sql.DB) *HabitRepository {
	return &HabitRepository{
		db: db,
	}
}

func (r *HabitRepository) GetAll() ([]*model.Habit, error) {
	return []*model.Habit{}, nil
}

func (r *HabitRepository) GetByID(id int) (*model.Habit, error) {
	return &model.Habit{}, nil
}

func (r *HabitRepository) Create(habit *model.Habit) (*model.Habit, error) {
	return &model.Habit{}, nil
}

func (r *HabitRepository) Update(id int, habit *model.Habit) (*model.Habit, error) {
	return &model.Habit{}, nil
}

func (r *HabitRepository) Delete(id int) error {
	return nil
}
