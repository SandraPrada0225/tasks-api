package repository

import (
	"database/sql"
	"fmt"
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

func (r *PostgresTaskRepository) Update(id int, title string) error {
	query := "UPDATE tasks SET title=$1 WHERE id=$2"

	result, err := r.DB.Exec(query, title, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("task no encontrada")
	}
	return nil
}

func (r *PostgresTaskRepository) Delete(id int) error {
	query := "DELETE FROM tasks WHERE id=$1"

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("task no encontrada")
	}
	return nil
}
