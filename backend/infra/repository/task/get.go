package task

import "backend/domain/model"

func (r *taskRepository) ListTasks() (tasks []*model.Task, err error) {
	err = r.DB.Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
