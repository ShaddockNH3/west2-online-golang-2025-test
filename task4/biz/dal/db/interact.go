package db

import (
	"errors"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/constants"
	"gorm.io/gorm"
)

type LikeItems struct {
	ID           string         `gorm:"primaryKey;type:varchar(100)"`
	UserID       string         // user_id，指的是点赞用户的

	LikeableID   string         `gorm:"index"` // 被点赞对象视频ID或评论ID
	LikeableType string         `gorm:"index"` // 被点赞对象的类型 "video" 或 "comment"

	CreatedAt    string         // create_at
	UpdatedAt    string         // update_at
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type CommentItems struct {
	ID         string `gorm:"primaryKey;type:varchar(100)"`
	UserId     string
	VideoId    string
	ParentId   string
	LikeCount  int64
	ChildCount int64
	Content    string
	CreateAt   string
	UpdateAt   string
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (LikeItems) TableName() string {
	return constants.LikesTableName
}

func (CommentItems) TableName() string {
	return constants.CommentsTableName
}

func CreateLike(likeableID, likeableType, userID string) error {
	if likeableType != "video" && likeableType != "comment" {
		return errors.New("invalid likeable type")
	}
	like := &LikeItems{
		UserID:       userID,
		LikeableID:   likeableID,
		LikeableType: likeableType,
	}
	return DB.Create(like).Error
}

func UpdateLike(likeableID, likeableType, userID string, likeType int64) error {
	if likeableType != "video" && likeableType != "comment" {
		return errors.New("invalid likeable type")
	}

	// 查询视频或者评论是否存在
	var exist bool
	var err error

	if likeableType == "video" {
		exist, err = IsVideoExist(likeableID)
	} else {
		exist, err = IsCommentExist(likeableID)
	}

	if err != nil {
		return err
	}

	if !exist {
		return errors.New("likeable item does not exist")
	}

	if likeType != 1 && likeType != 2 {
		return errors.New("invalid like type")
	}
	
	// 查询是否已经有记录
	var likeRecord LikeItems
	err = DB.Unscoped().Where("likeable_id = ? AND user_id = ?", likeableID, userID).First(&likeRecord).Error

	// 数据库里完全没有这条记录
	if err != nil && err == gorm.ErrRecordNotFound {
        if likeType == 1 { // 想点赞
            // 创建一个新的
            return CreateLike(likeableID, likeableType, userID)
        }
        if likeType == 2 { // 想取消赞
            // 本来就没有，不需要操作
            return errors.New("not liked yet")
        }
    }

	// 如果 err 不是 RecordNotFound，但还是有错误，那就直接返回
    if err != nil {
        return err
    }

	// 数据库里有记录，且没有被软删除
	if !likeRecord.DeletedAt.Valid { 
        if likeType == 1 { // 想点赞
            return errors.New("already liked")
        }
        if likeType == 2 { // 想取消赞
            // 软删除
            return DB.Where("id = ?", likeRecord.ID).Delete(&LikeItems{}).Error
        }
    }

	  // 是被“软删除”的记录
    if likeRecord.DeletedAt.Valid {
        if likeType == 1 {
            // 把 deleted_at 变回 null
            return DB.Unscoped().Model(&LikeItems{}).Where("id = ?", likeRecord.ID).Update("deleted_at", nil).Error
        }
        if likeType == 2 { // 想取消赞
            return errors.New("not liked yet")
        }
    }

	return nil
}

func QueryVideosByUserID(userID string, page, pageSize int64) ([]LikeItems, error) {
	var likes []LikeItems
	if err := DB.Where("user_id = ?", userID).Find(&likes).Error; err != nil {
		return nil, err
	}

	if err := DB.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&likes).Error; err != nil {
		return nil, err
	}

	return likes, nil
}

func QueryVideosByUserIDLike(videoID string) (VideoItems, error) {
	var video VideoItems
	if err := DB.Where("id = ?", videoID).First(&video).Error; err != nil {
		return VideoItems{}, err
	}
	return video, nil
}

func IsVideoExist(videoID string) (bool, error) {
	var count int64
	err := DB.Unscoped().Model(&VideoItems{}).Where("id = ?", videoID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func IsCommentExist(commentID string) (bool, error) {
	var count int64
	err := DB.Unscoped().Model(&CommentItems{}).Where("id = ?", commentID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
