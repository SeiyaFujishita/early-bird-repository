package active

import (
	"backend/domain/repository"
	"backend/infra"
)

type activeRepository struct {
	*infra.GormHandler
}

func NewActiveRepository(gormHandler *infra.GormHandler) repository.ActiveRepository {
	return &activeRepository{
		gormHandler,
	}
}
