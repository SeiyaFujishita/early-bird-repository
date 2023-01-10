package task

import (
	"backend/domain/repository"
	"backend/infra"
)

type taskRepository struct {
	*infra.GormHandler
}

func NewTaskRepository(gormHandler *infra.GormHandler) repository.TaskRepository {
	return &taskRepository{
		gormHandler,
	}
}
