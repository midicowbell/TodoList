package todo

import "time"

type Task struct {
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(title string, desc string) Task {
	return Task{
		Title:       title,
		Description: desc,
		IsCompleted: false,
		CreatedAt:   time.Now(),
	}
}

func (t *Task) Complete() {
	completeTime := time.Now()
	t.IsCompleted = true
	t.CompletedAt = &completeTime
}

func (t *Task) Uncomplete() {
	t.IsCompleted = false
	t.CompletedAt = nil
}
