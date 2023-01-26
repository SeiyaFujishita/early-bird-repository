package main

import (
	"backend/infra"
	"backend/server"
	"log"
	"time"

	domainrepo "backend/domain/repository"

	"backend/infra/repository/active"
	"backend/infra/repository/task"
	"backend/infra/repository/wakeup"

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
		TaskRepository:   task.NewTaskRepository(gormHandler),
		ActiveRepository: active.NewActiveRepository(gormHandler),
		WakeUpRepository: wakeup.NewWakeUpRepository(gormHandler),
	}

	// gin(フレームワーク)の初期化(https://gin-gonic.com/ja/)
	router := gin.Default()

	// CORS設定(https://www.tohoho-web.com/ex/cors.html)
	setCors(router)

	// usecase初期化処理
	server.NewApiServer(router, repos)

	log.Print("server start!")
	router.Run()
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
