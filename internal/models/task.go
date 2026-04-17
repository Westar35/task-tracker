package models

// Task представляет задачу с полями ID, Title и Done.
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Status    bool   `json:"status"`
	UserID    int64  `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CreateTaskRequest представляет структуру запроса для создания новой задачи.
type CreateTaskRequest struct {
	Title string `json:"title"`
}

type UpdateTaskRequest struct {
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
