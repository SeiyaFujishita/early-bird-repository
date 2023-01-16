package work

import (
	"backend/domain/repository"
	"backend/infra"
)

type userRepository struct {
	*infra.GormHandler
}

func NewUserRepository(gormHandler *infra.GormHandler) repository.UserRepository {
	return &userRepository{
		gormHandler,
	}
}
