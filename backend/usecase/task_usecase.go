package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"

	"github.com/gin-gonic/gin"
)

type TaskUsecase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		taskRepository: repo,
	}
}

func (u *TaskUsecase) ListTasks(c *gin.Context) ([]*model.Task, error) {
	return nil, nil
}
