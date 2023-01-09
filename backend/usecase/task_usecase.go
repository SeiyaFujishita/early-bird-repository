package usecase

import (
	"backend/domain/repository"
	"log"
	"net/http"

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

func (u *TaskUsecase) ListTasks(c *gin.Context) {
	tasks, err := u.taskRepository.ListTasks()

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}
