package repository

import "database/sql"

type HabitLogRepository struct {
	db *sql.DB
}

func NewHabitLogtRepository(db *sql.DB) *HabitLogRepository {
	return &HabitLogRepository{
		db: db,
	}
}