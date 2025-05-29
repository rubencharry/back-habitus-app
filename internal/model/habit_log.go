package model

type HabitLog struct {
	ID      int
	HabitID int
	Date string
	Completed string
}

type HabitLogDoc struct {
	ID      int `json:"id"`
	HabitID int `json:"habit_id"`
	Date string `json:"date"`
	Completed string `json:"completed"`
}

func HabitLogToDoc(habitLog *HabitLog) *HabitLogDoc {
	return &HabitLogDoc{
		ID:      habitLog.ID,
		HabitID: habitLog.HabitID,
		Date:    habitLog.Date,
		Completed: habitLog.Completed,
	}
}

func DocToHabitLog(habitLogDoc *HabitLogDoc) *HabitLog {
	return &HabitLog{
		ID:      habitLogDoc.ID,
		HabitID: habitLogDoc.HabitID,
		Date:    habitLogDoc.Date,
		Completed: habitLogDoc.Completed,
	}
}

func (h *HabitLog) UpdateModel(habitLogDoc *HabitLogDoc) {
	if habitLogDoc.HabitID != 0 {
		h.HabitID = habitLogDoc.HabitID
	}
	if habitLogDoc.Date != "" {
		h.Date = habitLogDoc.Date
	}
	if habitLogDoc.Completed != "" {
		h.Completed = habitLogDoc.Completed
	}
}