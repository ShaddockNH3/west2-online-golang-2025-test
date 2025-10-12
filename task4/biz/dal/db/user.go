package db

import (
	"time"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/configs/constants"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey;type:varchar(100)"`
	Username  string `gorm:"unique"`
	Password  string
	AvatarUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return constants.UserTableName
}

func CreateUser(user *User) error {
	return DB.Create(user).Error
}

func QueryUserByUsername(username string) (*User, error) {
	var user User
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func QueryUserByUserId(user_id string) (*User, error) {
	var user User
	if err := DB.Where("id = ?", user_id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UploadAvatar(user_id, avatar_url string) error {
	return DB.Model(&User{}).Where("id = ?", user_id).Update("avatar_url", avatar_url).Error
}