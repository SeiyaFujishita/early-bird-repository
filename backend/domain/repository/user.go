package repository

import "backend/domain/model"

type UserRepository interface {
	CreateUser(*model.User) error
}
