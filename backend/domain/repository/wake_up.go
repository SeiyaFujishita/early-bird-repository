package repository

import "backend/domain/model"

type WakeUpRepository interface {
	CreateWakeUp(*model.WakeUp) error
	GetWakeUpTime(int) (string, error)
}
