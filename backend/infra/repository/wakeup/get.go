package wakeup

import "backend/domain/model"

func (r *WakeUpRepository) GetWakeUpTime(id int) (string, error) {
	wakeUp := &model.WakeUp{}
	err := r.DB.Select("time").Where("user_id = ?", id).Find(wakeUp).Error

	if err != nil {
		return "", err
	}

	return wakeUp.Time, nil
}
