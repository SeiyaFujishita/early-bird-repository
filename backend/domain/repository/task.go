package repository

import "backend/domain/model"

type TaskRepository interface {
	ListTasks() ([]*model.Task, error)
}
