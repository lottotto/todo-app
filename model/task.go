package model

import "time"

type Task struct {
	Id int `json:"id"`

	UserId int `json:"user_id"`

	TypeId int `json:"type_id"`

	Title string `json:"title"`

	Detail string `json:"detail,omitempty"`

	Deadline time.Time `json:"deadline"`

	Done bool `json:"done"`
}
