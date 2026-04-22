package todo

import "time"

type Task struct {
	Title       string
	Description string
	IsComplited bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(title string, desc string) Task {
	return Task{
		Title:       title,
		Description: desc,
		IsComplited: false,
		CreatedAt:   time.Now(),
	}
}

func (t *Task) Complete() {
	completeTime := time.Now()
	t.IsComplited = true
	t.CompletedAt = &completeTime
}

func (t *Task) Uncomplete() {
	t.IsComplited = false
	t.CompletedAt = nil
}
