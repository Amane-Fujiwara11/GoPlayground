package usecase

import (
	"backend/domain"
	"backend/interface/repository"
)

type TaskUsecase struct {
	Repo repository.TaskRepository
}

func (u *TaskUsecase) GetAllTasks() ([]domain.Task, error) {
	return u.Repo.GetAll()
}

func (u *TaskUsecase) GetTaskByID(id int) (*domain.Task, error) {
	return u.Repo.GetByID(id)
}

func (u *TaskUsecase) CreateTask(task *domain.Task) error {
	return u.Repo.Create(task)
}

func (u *TaskUsecase) UpdateTask(task *domain.Task) error {
	return u.Repo.Update(task)
}

func (u *TaskUsecase) DeleteTask(id int) error {
	return u.Repo.Delete(id)
}
