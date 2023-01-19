package wakeup

import (
	"backend/domain/repository"
	"backend/infra"
)

type WakeUpRepository struct {
	*infra.GormHandler
}

func NewWakeUpRepository(gormHandler *infra.GormHandler) repository.WakeUpRepository {
	return &WakeUpRepository{
		gormHandler,
	}
}
