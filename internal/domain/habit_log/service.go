package habitlog

import "backend-habitus-app/internal/model"

type HabitLogService interface {
	GetAll() ([]*model.HabitLogDoc, error)
	Create(habitLog *model.HabitLogDoc) (*model.HabitLogDoc, error)
	Update(id int, habitLog *model.HabitLogDoc) (*model.HabitLogDoc, error)
}