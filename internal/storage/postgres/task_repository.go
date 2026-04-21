package postgres

import (
	"database/sql"
	"fmt"
	"task-tracker/internal/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAllTasks(userID int64) ([]models.Task, error) {
	// Создаем мапу задач. Ключ - ID задачи, значение - экземпляр структуры Task
	//tasks := make(map[int]models.TaskWithLocalNumber)
	tasks := make([]models.Task, 0)

	query := `
		SELECT
		id, title, status, user_id, created_at, updated_at
		FROM tasks
		WHERE user_id = $1
	`

	// Выполняем SQL-запрос для получения всех задач из базы данных
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Проходим по результатам запроса и заполняем мапу задач
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Status,
			&task.UserID,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepository) CreateTask(title string, userID int64) (models.Task, error) {

	var task models.Task

	status := false

	query := `
		INSERT INTO tasks (title, status, user_id)
		VALUES ($1, $2, $3)
		RETURNING id, title, status, user_id, created_at, updated_at;
	`

	err := r.db.QueryRow(query, title, status, userID).Scan(&task.ID, &task.Title, &task.Status, &task.UserID, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (r *TaskRepository) GetTaskByID(id int, userID int64) (models.Task, error) {
	var task models.Task

	query := `
		SELECT id, title, status
		FROM tasks
		WHERE id = $1 AND user_id = $2
	`

	err := r.db.QueryRow(query, id, userID).Scan(&task.ID, &task.Title, &task.Status)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepository) UpdateTaskByID(id int, title string, status bool) (models.Task, error) {
	var task models.Task

	query := `
		UPDATE tasks
		SET title = $1, status = $2
		WHERE id = $3
		RETURNING id, title, status;
	`

	err := r.db.QueryRow(query, title, status, id).Scan(&task.ID, &task.Title, &task.Status)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepository) DeleteTaskByID(id int) error {
	query := `
		DELETE FROM tasks
		WHERE id = $1
	`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
