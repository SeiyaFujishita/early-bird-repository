package work

import (
	"backend/domain/repository"
	"backend/infra"
)

type workRepository struct {
	*infra.GormHandler
}

func NewWorkRepository(gormHandler *infra.GormHandler) repository.WorkRepository {
	return &workRepository{
		gormHandler,
	}
}
