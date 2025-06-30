package models

import (
	"database/sql"
)

type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

func GetTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("SELECT id, title, content, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Content, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func CreateTask(db *sql.DB, task *Task) error {
	task.Status = "registered"
	result, err := db.Exec("INSERT INTO tasks (title, content, status) VALUES (?, ?, ?)", task.Title, task.Content, task.Status)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	task.ID = int(id)
	return nil
}

func DeleteTask(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

func UpdateTaskStatus(db *sql.DB, id int, status string) error {
	_, err := db.Exec("UPDATE tasks SET status = ? WHERE id = ?", status, id)
	return err
}
