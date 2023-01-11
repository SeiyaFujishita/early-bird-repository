package server

import (
	"backend/domain/repository"
	"backend/usecase"

	"github.com/gin-gonic/gin"
)

func NewApiServer(router *gin.Engine, repos repository.Repositories) {
	taskUsecase := usecase.NewTaskUsecase(repos.TaskRepository)
	workUsecase := usecase.NewWorkUsecase(repos.WorkRepository)

	router.GET("/tasks", func(c *gin.Context) { taskUsecase.ListTasks(c) })
	router.POST("/work", func(c *gin.Context) { workUsecase.CreateWork(c) })
}
