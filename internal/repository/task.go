package repository

import "database/sql"

type TaskRepository struct {
	db *sql.DB
}

func NewTasktRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}