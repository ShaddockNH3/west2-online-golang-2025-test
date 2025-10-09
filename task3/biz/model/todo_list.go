package model

import "gorm.io/gorm"

type ToDoList struct {
	gorm.Model
	UserID  int64  `json:"user_id" gorm:"column:user_id"`
	Title   string `json:"title" gorm:"column:title"`
	Context string `json:"context" gorm:"column:context"`
	Status  int    `json:"status" gorm:"column:status"` 
}

func(t *ToDoList) TableName()string{
	return "todo_list"
}