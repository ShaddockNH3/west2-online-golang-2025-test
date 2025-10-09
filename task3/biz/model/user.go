package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name" gorm:"column:name"`
	Introduce string `json:"introduce" gorm:"column:introduce"`
	Password  string `json:"password" gorm:"column:password"`
}

func (u *User) TableName() string {
	return "users"
}
