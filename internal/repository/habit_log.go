package repository

import (
	"backend-habitus-app/internal/model"
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

type HabitLogRepository struct {
	db *sql.DB
}

func NewHabitLogtRepository(db *sql.DB) *HabitLogRepository {
	return &HabitLogRepository{
		db: db,
	}
}

func (r *HabitLogRepository) GetAll() ([]*model.HabitLog, error) {
	rows, err := r.db.Query("SELECT id, habit_id, date, completed FROM habit_logs")

	if err != nil {
		if err == sql.ErrNoRows {
			return []*model.HabitLog{}, nil
		}
		return nil, errors.New("internal server error")
	}

	habits := make([]*model.HabitLog, 0)
	for rows.Next() {
		var habitLog model.HabitLog
		if err := rows.Scan(&habitLog.ID, &habitLog.HabitID, &habitLog.Date, &habitLog.Completed); err != nil {
			return nil, errors.New("internal server error")
		}
		habits = append(habits, &habitLog)
	}

	return habits, nil
}

func (r *HabitLogRepository) GetByID(id int) (*model.HabitLog, error) {
row := r.db.QueryRow("SELECT id, habit_id, date, completed FROM habit_logs WHERE id = ?", id)
	if err := row.Err(); err != nil {
		return nil, errors.New("internal server error")
	}

	var habitLog model.HabitLog
	if err := row.Scan(&habitLog.ID, &habitLog.HabitID, &habitLog.Date, &habitLog.Completed); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("habit log not found")
		}
		return nil, errors.New("internal server error")
	}
	
	return &habitLog, nil
}

func (r *HabitLogRepository) Create(habitLog *model.HabitLog) (*model.HabitLog, error) {
	tx, err := r.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err != nil {
		return nil, errors.New("internal server error")
	}

	result, err := r.db.Exec("INSERT INTO habit_logs (habit_id, date, completed) VALUES (?, ?, ?)",
		habitLog.HabitID, habitLog.Date, habitLog.Completed)
	if err != nil {
		if sqlErr, ok := err.(*mysql.MySQLError); ok {
			switch sqlErr.Number {
			case 1406:
				return nil, errors.New("habit log format not accepted")
			case 1452:
				return nil, errors.New("habit_id does not exist")
			}
		}

		return nil, errors.New("internal server error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.New("internal server error")
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New("internal server error")
	}

	habitLog.ID = int(id)
	return habitLog, nil
}

func (r *HabitLogRepository) Update(id int,habitLog *model.HabitLog) (*model.HabitLog, error) {

	tx, err := r.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err != nil {
		return nil, errors.New("internal server error")
	}

	result, err := r.db.Exec("UPDATE habit_logs SET habit_id = ?, date = ?, completed = ? WHERE id = ?",
		habitLog.HabitID, habitLog.Date, habitLog.Completed, id)
	if err != nil {
		if sqlErr, ok := err.(*mysql.MySQLError); ok {
			switch sqlErr.Number {
			case 1406:
				return nil, errors.New("habit log format not accepted")
			case 1452:
				return nil, errors.New("habit_id does not exist")
			}
		}

		return nil, errors.New("internal server error")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return nil, errors.New("habit log not found")
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return habitLog, nil
}
