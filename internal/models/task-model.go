package models

// Task представляет задачу с полями ID, Title и Done.
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// CreateTaskRequest представляет структуру запроса для создания новой задачи.
type CreateTaskRequest struct {
	Title string `json:"title"`
}
