package mysql

import (
	"time"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task5/pkg/constants"
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
	MfaSecret string         `gorm:"type:varchar(255)"`
}

func (User) TableName() string {
	return constants.UserTableName
}

type Image struct {
	ID               string `gorm:"primaryKey;type:varchar(100)"`
	UserID           string
	URL              string
	OriginalFilename string
	Filepath         string
	Filesize         int64
	MimeType         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

func (Image) TableName() string {
	return constants.ImageTableName 
}
