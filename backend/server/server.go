package server

import (
	"backend/domain/repository"
	"backend/usecase"

	"github.com/gin-gonic/gin"
)

func NewApiServer(router *gin.Engine, repos repository.Repositories) {
	taskUsecase := usecase.NewTaskUsecase(repos.TaskRepository)
	activeUsecase := usecase.NewActiveUsecase(repos.ActiveRepository)
	userUsecase := usecase.NewUserUsecase(repos.UserRepository)

	router.GET("/tasks", func(c *gin.Context) { taskUsecase.ListTasks(c) })
	router.POST("/actives", func(c *gin.Context) { activeUsecase.CreateActive(c) })
	router.POST("/user", func(c *gin.Context) { userUsecase.CreateUser(c) })
}
