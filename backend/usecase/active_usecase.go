package usecase

import (
	"fmt"
	"backend/domain/model"
	"backend/domain/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ActiveUsecase struct {
	activeRepository repository.ActiveRepository
}

func NewActiveUsecase(repo repository.ActiveRepository) *ActiveUsecase {
	return &ActiveUsecase{
		activeRepository: repo,
	}
}

func (u *ActiveUsecase) CreateActive(c *gin.Context) {
	active := &model.Active{}

	err := c.ShouldBindWith(active, binding.JSON)
	fmt.Println("------------------------")
	fmt.Printf("%+v\n", err)
	fmt.Println("------------------------")
	if err != nil {
		log.Print(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = u.activeRepository.CreateActive(active)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
