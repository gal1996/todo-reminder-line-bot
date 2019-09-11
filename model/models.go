package model

import "time"

type Task struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Content string `json:"content"`
	DeadLine time.Time `json:"dead_line"`
	IsDone bool `json:"is_done"`
}

type User struct {
	Id int `json:"id"`
	UserName string `json:"user_name"`
}
