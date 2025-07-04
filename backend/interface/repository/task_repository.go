package repository

import "backend/domain"

type TaskRepository interface {
	GetAll() ([]domain.Task, error)
	GetByID(id int) (*domain.Task, error)
	Create(task *domain.Task) error
	Update(task *domain.Task) error
	Delete(id int) error
}
