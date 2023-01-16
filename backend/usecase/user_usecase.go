package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: repo,
	}
}

func (u *UserUsecase) CreateUser(c *gin.Context) {
	user := &model.User{}

	err := c.ShouldBindWith(user, binding.JSON)

	if err != nil {
		log.Print(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = u.userRepository.CreateUser(user)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
