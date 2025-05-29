package service

import (
	habitlog "backend-habitus-app/internal/domain/habit_log"
	"backend-habitus-app/internal/model"
	"errors"
)

type HabitLogService struct {
	rp habitlog.HabitLogRepository
}

func NewHabitLogService(rp habitlog.HabitLogRepository) *HabitLogService {
	return &HabitLogService{
		rp: rp,
	}
}

func (s *HabitLogService) GetAll() ([]*model.HabitLogDoc, error) {
	
	habits, err := s.rp.GetAll()
	if err != nil {
		return nil, err
	}

	var habitlogs []*model.HabitLogDoc
	for _, h := range habits {
		habitLog := model.HabitLogToDoc(h)
		habitlogs = append(habitlogs, habitLog)
	}

	return habitlogs, nil
}

func (s *HabitLogService) Create(habitLog *model.HabitLogDoc) (*model.HabitLogDoc, error) {
	
	habitLogToModel := model.DocToHabitLog(habitLog)
	createdHabitLog, err := s.rp.Create(habitLogToModel)
	if err != nil {
		return nil, err
	}
	return model.HabitLogToDoc(createdHabitLog), nil
}

func (s *HabitLogService) Update(id int, habitLog *model.HabitLogDoc) (*model.HabitLogDoc, error) {
	if id < 0 {
		return nil, errors.New("invalid habit log ID")
	}
	
	habitLogModel, err := s.rp.GetByID(id)
	if err != nil {
		return nil, err
	}

	habitLogModel.UpdateModel(habitLog)
	habitLogModel.ID = id

	_,err = s.rp.Update(id, habitLogModel)
	if err != nil {
		return nil, err
	}

	*habitLog = *model.HabitLogToDoc(habitLogModel)
	return habitLog, nil
}