package mysql

import (
	"errors"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/model"
	"gorm.io/gorm"
)

func CreateUser(users *model.User) error {
	return DB.Create(users).Error
}

func DeleteUser(userId int64) error {
	return DB.Where("id=?", userId).Delete(&model.User{}).Error
}

func UpdateUser(user *model.User) error {
	return DB.Where("id = ?", user.ID).Updates(user).Error
}

func QueryUser(keyword *string, page, pageSize int64) ([]*model.User, int64, error) {
	db := DB.Model(model.User{})
	if keyword != nil && len(*keyword) != 0 {
		db = db.Where(DB.Or("name like ?", "%"+*keyword+"%").
			Or("introduce like ?", "%"+*keyword+"%"))
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []*model.User
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func QueryUserByName(name string) (*model.User, error) {
	var user model.User

	err := DB.Where("name = ?", name).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
