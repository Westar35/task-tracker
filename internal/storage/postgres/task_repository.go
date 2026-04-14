package postgres

import (
	"database/sql"
	"task-tracker/internal/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAllTasks() (map[int]models.Task, error) {
	// Создаем мапу задач. Ключ - ID задачи, значение - экземпляр структуры Task
	tasks := make(map[int]models.Task)

	// Выполняем SQL-запрос для получения всех задач из базы данных
	rows, err := r.db.Query("SELECT id, title, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Проходим по результатам запроса и заполняем мапу задач
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Status)
		if err != nil {
			return nil, err
		}
		tasks[task.ID] = task
	}
	return tasks, nil
}

func (r *TaskRepository) CreateTask(title string) (models.Task, error) {

	var task models.Task

	query := `
		INSERT INTO tasks (title, status)
		VALUES ($1, false)
		RETURNING id, title, status
	`

	err := r.db.QueryRow(query, title).Scan(&task.ID, &task.Title, &task.Status)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}
