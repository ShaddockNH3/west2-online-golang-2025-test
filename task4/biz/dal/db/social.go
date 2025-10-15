package db

import (
	"errors"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/common"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/constants"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SocialItems struct {
	ID         string `gorm:"primaryKey;type:varchar(100)"`
	FollowerID string // 关注者ID
	FollowedID string // 被关注者ID
	CreatedAt  string
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (SocialItems) TableName() string {
	return constants.FollowsTableName
}

func CreateSocialItems(items *SocialItems) error {
	return DB.Create(items).Error
}

func UpdateRelation(follower_id, followed_id string, action_type int64) error {
	var err error
	var social_relation SocialItems
	// 先判断是否是合法操作
	if action_type != 0 && action_type != 1 {
		return errors.New("action_type error")
	}

	// 0-关注 1-取消关注
	err = DB.Unscoped().Where("follower_id = ? AND followed_id = ?", follower_id, followed_id).First(&social_relation).Error

	// 数据库里完全没有这条记录
	if err != nil && err == gorm.ErrRecordNotFound {
		if action_type == 0 { // 想关注
			// 创建一个新的
			return CreateSocialItems(&SocialItems{
				ID:         uuid.New().String(),
				FollowerID: follower_id,
				FollowedID: followed_id,
			})
		}
		if action_type == 1 { // 想取消关注
			// 本来就没有，不需要操作
			return errors.New("not followed yet")
		}
	}

	// 如果 err 不是 RecordNotFound，但还是有错误，那就直接返回
	if err != nil {
		return err
	}

	// 数据库里有记录，且没有被软删除
	if !social_relation.DeletedAt.Valid {
		if action_type == 0 { // 想关注
			return errors.New("already followed")
		}
		if action_type == 1 { // 想取消关注
			// 软删除
			return DB.Where("id = ?", social_relation.ID).Delete(&SocialItems{}).Error
		}
	}

	// 是被“软删除”的记录
	if social_relation.DeletedAt.Valid {
		if action_type == 0 {
			// 把 deleted_at 变回 null
			return DB.Unscoped().Model(&SocialItems{}).Where("id = ?", social_relation.ID).Update("deleted_at", nil).Error
		}
		if action_type == 1 { // 想取消关注
			return errors.New("not followed yet")
		}
	}

	return nil
}

func QueryFollowingList(userID string, page, pageSize int64) ([]*common.SocialDTO, error) {
	var follows []SocialItems
	if err := DB.Where("follower_id = ?", userID).Find(&follows).Error; err != nil {
		return nil, err
	}

	var followed []string
	for _, follow := range follows {
		followed = append(followed, follow.FollowedID)
	}

	var users []common.SocialDTO
	err := DB.Model(&User{}).
		Where("id IN ?", followed).
		Limit(int(pageSize)).
		Offset(int(pageSize * (page - 1))).
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	var userPtrs []*common.SocialDTO
	for i := range users {
		userPtrs = append(userPtrs, &users[i])
	}
	return userPtrs, nil
}

func QueryFollowerList(userID string, page, pageSize int64) ([]*common.SocialDTO, error) {
	var follows []SocialItems
	if err := DB.Where("followed_id = ?", userID).Find(&follows).Error; err != nil {
		return nil, err
	}

	var followers []string
	for _, follow := range follows {
		followers = append(followers, follow.FollowerID)
	}

	var users []common.SocialDTO
	err := DB.Model(&User{}).
		Where("id IN ?", followers).
		Limit(int(pageSize)).
		Offset(int(pageSize * (page - 1))).
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	var userPtrs []*common.SocialDTO
	for i := range users {
		userPtrs = append(userPtrs, &users[i])
	}
	return userPtrs, nil
}

func QueryFriendList(myID string, page, pageSize int64) ([]*common.SocialDTO, error) {
	var friendIDs []string

	// t1 代表 我关注TA 这条记录
	// t2 代表 TA关注我 这条记录
	// 要找到同时满足这两个条件的记录
	err := DB.Table("follows as t1").
		Select("t1.followed_id").
		Joins("JOIN follows as t2 ON t1.follower_id = t2.followed_id AND t1.followed_id = t2.follower_id").
		Where("t1.follower_id = ?", myID).
		Scan(&friendIDs).Error

	if err != nil {
		return nil, err
	}

	if len(friendIDs) == 0 {
		return nil, errors.New("no friends found")
	}

	var users []common.SocialDTO
	err = DB.Model(&User{}).
		Where("id IN ?", friendIDs).
		Limit(int(pageSize)).
		Offset(int(pageSize * (page - 1))).
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	var userPtrs []*common.SocialDTO
	for i := range users {
		userPtrs = append(userPtrs, &users[i])
	}
	return userPtrs, nil
}
