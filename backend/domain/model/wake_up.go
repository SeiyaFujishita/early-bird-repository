package model

type WakeUp struct {
		Id     int `json:"id" `
		Time   string `json:"time" binding:"required" `
		Name string `json:"name" binding:"required"`
	}
