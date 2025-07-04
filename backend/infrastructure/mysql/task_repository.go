package mysql

import (
	"backend/domain"
	"backend/interface/repository"
	"database/sql"
)

type TaskRepositoryImpl struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) repository.TaskRepository {
	return &TaskRepositoryImpl{DB: db}
}

func (r *TaskRepositoryImpl) GetAll() ([]domain.Task, error) {
	rows, err := r.DB.Query("SELECT id, title, content, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Content, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepositoryImpl) GetByID(id int) (*domain.Task, error) {
	row := r.DB.QueryRow("SELECT id, title, content, status FROM tasks WHERE id = ?", id)
	var task domain.Task
	if err := row.Scan(&task.ID, &task.Title, &task.Content, &task.Status); err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepositoryImpl) Create(task *domain.Task) error {
	task.Status = "registered"
	result, err := r.DB.Exec("INSERT INTO tasks (title, content, status) VALUES (?, ?, ?)", task.Title, task.Content, task.Status)
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

func (r *TaskRepositoryImpl) Update(task *domain.Task) error {
	_, err := r.DB.Exec("UPDATE tasks SET title = ?, content = ?, status = ? WHERE id = ?", task.Title, task.Content, task.Status, task.ID)
	return err
}

func (r *TaskRepositoryImpl) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
