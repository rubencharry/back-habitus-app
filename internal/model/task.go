package model

type Task struct {
	ID    int
	Title string
	Description string
	DueDate string
	Completed bool
	Created_at string
}

type TaskDoc struct {
	ID    int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	DueDate string `json:"due_date"`
	Completed bool `json:"completed"`
	Created_at string `json:"created_at"`
}

func TaskToDoc(task *Task) *TaskDoc {
	return &TaskDoc{
		ID:    task.ID,
		Title: task.Title,
		Description: task.Description,
		DueDate: task.DueDate,
		Completed: task.Completed,
		Created_at: task.Created_at,
	}
}

func DocToTask(taskDoc *TaskDoc) *Task {
	return &Task{
		ID:    taskDoc.ID,
		Title: taskDoc.Title,
		Description: taskDoc.Description,
		DueDate: taskDoc.DueDate,
		Completed: taskDoc.Completed,
		Created_at: taskDoc.Created_at,
	}
}

func (t *Task) UpdateModel(taskDoc *TaskDoc) {
	if taskDoc.Title != "" {
		t.Title = taskDoc.Title
	}
	if taskDoc.Description != "" {
		t.Description = taskDoc.Description
	}
	if taskDoc.DueDate != "" {
		t.DueDate = taskDoc.DueDate
	}
	if taskDoc.Completed {
		t.Completed = taskDoc.Completed
	}
	if taskDoc.Created_at != "" {
		t.Created_at = taskDoc.Created_at
	}
}