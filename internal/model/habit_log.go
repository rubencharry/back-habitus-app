package model

type HabitLog struct {
	ID      int
	HabitID int
	Date string
	Completed string
	CreatedAt string
}

type HabitLogDoc struct {
	ID      int `json:"id"`
	HabitID int `json:"habit_id"`
	Date string `json:"date"`
	Completed string `json:"completed"`
	CreatedAt string `json:"created_at"`
}

func HabitLogToDoc(habitLog *HabitLog) *HabitLogDoc {
	return &HabitLogDoc{
		ID:      habitLog.ID,
		HabitID: habitLog.HabitID,
		Date:    habitLog.Date,
		Completed: habitLog.Completed,
		CreatedAt: habitLog.CreatedAt,
	}
}

func DocToHabitLog(habitLogDoc *HabitLogDoc) *HabitLog {
	return &HabitLog{
		ID:      habitLogDoc.ID,
		HabitID: habitLogDoc.HabitID,
		Date:    habitLogDoc.Date,
		Completed: habitLogDoc.Completed,
		CreatedAt: habitLogDoc.CreatedAt,
	}
}