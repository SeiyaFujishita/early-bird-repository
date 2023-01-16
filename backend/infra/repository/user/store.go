package work

import "backend/domain/model"

func (r *userRepository) CreateUser(user *model.User) (err error) {
	err = r.DB.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}
