package repository

import (
	"database/sql"
	"tasks-api/internal/models"
)

// implementacion de postgres
type PostgresTaskRepository struct {
	DB *sql.DB
}

func (r *PostgresTaskRepository) GetAll() ([]models.Task, error) {
	rows, err := r.DB.Query("SELECT id, title FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.ID, &t.Title)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *PostgresTaskRepository) GetByID(id int) (models.Task, error) {
	var task models.Task

	query := "SELECT id, title FROM tasks WHERE id=$1"

	err := r.DB.QueryRow(query, id).Scan(&task.ID, &task.Title)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *PostgresTaskRepository) Create(title string) (models.Task, error) {
	var task models.Task
	query := "INSERT INTO tasks (title) VALUES ($1) RETURNING id"
	err := r.DB.QueryRow(query, title).Scan(&task.ID)
	if err != nil {
		return task, err
	}
	task.Title = title
	return task, nil
}

func (r *PostgresTaskRepository) Update(id int, title string) (int64, error) {
	query := "UPDATE tasks SET title=$1 WHERE id=$2"

	result, err := r.DB.Exec(query, title, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (r *PostgresTaskRepository) Delete(id int) (int64, error) {
	query := "DELETE FROM tasks WHERE id=$1"

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
