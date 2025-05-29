package repository

import (
	"backend-habitus-app/internal/model"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTasktRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) GetAll() ([]*model.Task, error) {
	rows, err := r.db.Query("SELECT id, title, description, due_date, completed, created_at FROM tasks")

	if err != nil {
		if err == sql.ErrNoRows {
			return []*model.Task{}, nil
		}
		return nil, errors.New("internal server error")
	}

	tasks := make([]*model.Task, 0)
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed, &task.Created_at); err != nil {
			return nil, errors.New("internal server error")
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (r *TaskRepository) GetByID(id int) (*model.Task, error) {
	row := r.db.QueryRow("SELECT id, title, description, due_date, completed, created_at FROM tasks WHERE id = ?", id)

	if err := row.Err(); err != nil {
		return nil, errors.New("internal server error")
	}

	var task model.Task
	if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed, &task.Created_at); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("task not found")
		}
		return nil, errors.New("internal server error")
	}

	return &task, nil
}

func (r *TaskRepository) Create(task *model.Task) (*model.Task, error) {

	tx, err := r.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err != nil {
		return nil, errors.New("internal server error")
	}

	result, err := tx.Exec("INSERT INTO tasks (title, description, due_date, completed) VALUES (?, ?, ?, ?)", task.Title, task.Description, task.DueDate, task.Completed)

	if err != nil {
		if sqlErr, ok := err.(*mysql.MySQLError); ok {
			switch sqlErr.Number {
			case 1062:
				return nil, errors.New("task already exists")
			case 1406:
				return nil, errors.New("task format not acepted")
			}
		}
		return nil, errors.New("internal server error")
	}

	rows, err := result.RowsAffected()
	if err != nil || rows != 1 {
		return nil, errors.New("internal server error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.New("internal server error")
	}

	var createdAt string
	err = tx.QueryRow("SELECT created_at FROM tasks WHERE id = ?", id).Scan(&createdAt)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New("internal server error")
	}
	task.ID = int(id)
	task.Created_at = createdAt

	return task, nil
}

func (r *TaskRepository) Update(id int, task *model.Task) (*model.Task, error) {

	tx, err := r.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err != nil {
		return nil, errors.New("internal server error")
	}

	result, err := tx.Exec("UPDATE tasks SET title = ?, description = ?, due_date = ?, completed = ? WHERE id = ?", task.Title, task.Description, task.DueDate, task.Completed, id)
	if err != nil {
		if sqlErr, ok := err.(*mysql.MySQLError); ok {
			switch sqlErr.Number {
			case 1062:
				return nil, errors.New("task already exists")
			case 1406:
				return nil, errors.New("task format not accepted")
			}
		}
		return nil, errors.New("internal server error")
	}

	rows, err := result.RowsAffected()
	if err != nil || rows > 1 {
		tx.Rollback()
		return nil, errors.New("internal server error")
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return task, nil
}

func (r *TaskRepository) Delete(id int) error {
	tx, err := r.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err != nil {
		return errors.New("internal server error")
	}

	result, err := tx.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return errors.New("internal server error")
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return errors.New("task not found")
	}

	if rows > 1 {
		return errors.New("internal server error")
	}

	err = tx.Commit()
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}
