package active

import "backend/domain/model"

func (r *activeRepository) CreateActive(active *model.Active) (err error) {
	err = r.DB.Create(&active).Error

	if err != nil {
		return err
	}

	return nil
}
