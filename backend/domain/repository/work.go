package repository

import "backend/domain/model"

type WorkRepository interface {
	CreateWork(*model.Work) error
}
