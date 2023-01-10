package server

import (
	"backend/domain/repository"
	"backend/usecase"

	"github.com/gin-gonic/gin"
)

func NewApiServer(router *gin.Engine, repos repository.Repositories) {
	taskUsecase := usecase.NewTaskUsecase(repos.TaskRepository)

	// GithubAPIによるデータ取得
	router.GET("/tasks", func(c *gin.Context) { taskUsecase.ListTasks(c) })
}
