package model

type Work struct {
	Id     int    `json:"id" `
	Time   string `json:"time" binding:"required" `
	TaskId int    `json:"taskId" binding:"required"`
}
