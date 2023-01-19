package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type WakeUpUsecase struct {
	WakeUpRepository repository.WakeUpRepository
}

func NewWakeUpUsecase(repo repository.WakeUpRepository) *WakeUpUsecase {
	return &WakeUpUsecase{
		WakeUpRepository: repo,
	}
}

func (u *WakeUpUsecase) CreateWakeUp(c *gin.Context) {
	wakeUp := &model.WakeUp{}

	err := c.ShouldBindWith(wakeUp, binding.JSON)
	fmt.Println("------------------------")
	fmt.Printf("%+v\n", wakeUp)
	fmt.Println("------------------------")
	if err != nil {
		log.Print(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("------------------------")
	fmt.Printf("%+v\n", wakeUp)
	fmt.Println("------------------------")
	err = u.WakeUpRepository.CreateWakeUp(wakeUp)
	fmt.Println("------------------------")
	fmt.Printf("%+v\n", err)
	fmt.Println("------------------------")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (u *WakeUpUsecase) GetWakeUpTime(c *gin.Context) {
	id, err := strconv.Atoi((c.Param("user_id")))

	if err != nil {
		log.Print(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	time, err := u.WakeUpRepository.GetWakeUpTime(id)

	if err != nil {
		log.Print(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, time)
}
