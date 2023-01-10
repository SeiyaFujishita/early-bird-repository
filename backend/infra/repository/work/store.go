package work

import "backend/domain/model"

func (r *workRepository) CreateWork(work *model.Work) (err error) {
	err = r.DB.Create(&work).Error

	if err != nil {
		return err
	}

	return nil
}
