package model

import "time"

// Task описывает структуру задачи в нашей системе
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // "todo", "in_progress", "done"
	CreatedAt   time.Time `json:"created_at"`
}
