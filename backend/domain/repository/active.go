package repository

import "backend/domain/model"

type ActiveRepository interface {
	CreateActive(*model.Active) error
}
