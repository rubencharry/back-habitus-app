package repository

import (
	"backend-habitus-app/internal/model"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
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
	rows, err := r.db.Query("SELECT id, title, description, frequency, created_at FROM habits")
	
	if err != nil {
		if err == sql.ErrNoRows {
			return []*model.Habit{}, nil
		}
		return nil, errors.New("internal server error")
	}

	habits := make([]*model.Habit, 0)
	for rows.Next() {
		var habit model.Habit
		if err := rows.Scan(&habit.ID, &habit.Title, &habit.Description, &habit.Frequency, &habit.CreatedAt); err != nil {
			return nil, errors.New("internal server error")
		}
		habits = append(habits, &habit)
	}

	return habits, nil
}

func (r *HabitRepository) GetByID(id int) (*model.Habit, error) {
	row := r.db.QueryRow("SELECT id, title, description, frequency, created_at FROM habits WHERE id = ?", id)

	if err := row.Err(); err != nil {
		return nil, errors.New("internal server error")
	}

	var habit model.Habit
	if err := row.Scan(&habit.ID, &habit.Title, &habit.Description, &habit.Frequency, &habit.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("habit not found")
		}
		return nil, errors.New("internal server error")
	}

	return &habit, nil
}

func (r *HabitRepository) Create(habit *model.Habit) (*model.Habit, error) {
	tx, err := r.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err != nil {
		return nil, errors.New("internal server error")
	}

	result, err := tx.Exec("INSERT INTO habits (title, description, frequency) VALUES (?, ?, ?)", habit.Title, habit.Description, habit.Frequency)
	if err != nil {
		if sqlErr, ok := err.(*mysql.MySQLError); ok {
			switch sqlErr.Number {
			case 1062:
				return nil, errors.New("habit already exists")
			case 1406:
				return nil, errors.New("habit format not accepted")
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
	err = tx.QueryRow("SELECT created_at FROM habits WHERE id = ?", id).Scan(&createdAt)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New("internal server error")
	}

	habit.ID = int(id)
	habit.CreatedAt = createdAt

	return habit, nil
}


func (r *HabitRepository) Update(id int, habit *model.Habit) (*model.Habit, error) {

	tx, err := r.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err != nil {
		return nil, errors.New("internal server error")
	}

	result, err := tx.Exec("UPDATE habits SET title = ?, description = ?, frequency = ? WHERE id = ?", habit.Title, habit.Description, habit.Frequency, id)
	if err != nil {
		if sqlErr, ok := err.(*mysql.MySQLError); ok {
			switch sqlErr.Number {
			case 1062:
				return nil, errors.New("habit already exists")
			case 1406:
				return nil, errors.New("habit format not accepted")
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

	return habit, nil
}

func (r *HabitRepository) Delete(id int) error {
	tx, err := r.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err != nil {
		return errors.New("internal server error")
	}

	result, err := tx.Exec("DELETE FROM habits WHERE id = ?", id)
	if err != nil {
		return errors.New("internal server error")
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return errors.New("habit not found")
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
