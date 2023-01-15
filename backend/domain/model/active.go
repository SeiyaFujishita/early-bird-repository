package model

type Active struct {
	Id     int `json:"id" `
	Time   int `json:"time" binding:"required" `
	TaskId int `json:"taskId" binding:"required"`
	UserId int `json:"userId" binding:"required"`
}
