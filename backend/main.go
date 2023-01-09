package main

import (
	"backend/infra"
	"backend/server"
	"log"
	"time"

	domainrepo "backend/domain/repository"

	"backend/infra/repository/task"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 環境変数読み込み
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Failed to connect env: %v", err)
	}

	// DB初期化処理
	gormHandler := infra.NewGormHandler()

	// リポジトリ初期化処理
	repos := domainrepo.Repositories{
		TaskRepository: task.NewTaskRepository(gormHandler),
	}

	router := gin.Default()

	setCors(router)

	// usecase初期化処理
	server.NewApiServer(router, repos)

	router.Run()
	log.Print("server start!")
}

func setCors(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
