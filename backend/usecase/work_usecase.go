package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type WorkUsecase struct {
	workRepository repository.WorkRepository
}

func NewWorkUsecase(repo repository.WorkRepository) *WorkUsecase {
	return &WorkUsecase{
		workRepository: repo,
	}
}

func (u *WorkUsecase) CreateWork(c *gin.Context) {
	work := &model.Work{}

	err := c.ShouldBindWith(work, binding.JSON)

	if err != nil {
		log.Print(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = u.workRepository.CreateWork(work)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
