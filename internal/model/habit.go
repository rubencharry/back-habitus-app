package model

type Habit struct {
	ID    int
	Title string
	Description string
	Frequency string 
	CreatedAt string
}

type HabitDoc struct {
	ID    int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Frequency string `json:"frequency"`
	CreatedAt string `json:"created_at"`
}

func HabitToDoc(habit *Habit) *HabitDoc {
	return &HabitDoc{
		ID:    habit.ID,
		Title: habit.Title,
		Description: habit.Description,
		Frequency: habit.Frequency,
		CreatedAt: habit.CreatedAt,
	}
}

func DocToHabit(habitDoc *HabitDoc) *Habit {
	return &Habit{
		ID:    habitDoc.ID,
		Title: habitDoc.Title,
		Description: habitDoc.Description,
		Frequency: habitDoc.Frequency,
		CreatedAt: habitDoc.CreatedAt,
	}
}