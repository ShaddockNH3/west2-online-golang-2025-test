package db

import (
	"errors"
	"time"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/common"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/constants"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LikeItems struct {
	ID     string `gorm:"primaryKey;type:varchar(100)"`
	UserID string // user_id，指的是点赞用户的

	LikeableID   string `gorm:"index"` // 被点赞对象视频ID或评论ID
	LikeableType string `gorm:"index"` // 被点赞对象的类型 "video" 或 "comment"

	CreatedAt time.Time      // create_at
	UpdatedAt time.Time      // update_at
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CommentItems struct {
	ID         string `gorm:"primaryKey;type:varchar(100)"`
	UserId     string
	VideoId    string
	ParentId   string // 父评论ID，若为一级评论则为空
	LikeCount  int64
	ChildCount int64
	Content    string
	CreateAt   time.Time
	UpdateAt   time.Time
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
		ID:           uuid.New().String(),
		UserID:       userID,
		LikeableID:   likeableID,
		LikeableType: likeableType,
	}
	return DB.Create(like).Error
}

func CreateComment(comment *CommentItems) error {
	return DB.Create(comment).Error
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
			err = CreateLike(likeableID, likeableType, userID)
			if err != nil {
				return err
			}
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
			err = DB.Where("id = ?", likeRecord.ID).Delete(&LikeItems{}).Error
			if err != nil {
				return err
			}
		}
	}

	// 是被“软删除”的记录
	if likeRecord.DeletedAt.Valid {
		if likeType == 1 {
			// 把 deleted_at 变回 null
			err = DB.Unscoped().Model(&LikeItems{}).Where("id = ?", likeRecord.ID).Update("deleted_at", nil).Error
			if err != nil {
				return err
			}
		}
		if likeType == 2 { // 想取消赞
			return errors.New("not liked yet")
		}
	}

	// 更新点赞数
	if likeableType == "video" {
		video, err := GetVideosByIDLike(likeableID)
		if err != nil {
			return err
		}
		var newLikeCount int64
		if likeType == 1 {
			newLikeCount = video.LikeCount + 1
		} else {
			newLikeCount = video.LikeCount - 1
		}
		if newLikeCount < 0 {
			newLikeCount = 0
		}
		if err := DB.Model(&VideoItems{}).Where("id = ?", video.ID).Update("like_count", newLikeCount).Error; err != nil {
			return err
		}
	} else {
		comment, err := GetCommentsByIDLike(likeableID)
		if err != nil {
			return err
		}
		var newLikeCount int64
		if likeType == 1 {
			newLikeCount = comment.LikeCount + 1
		} else {
			newLikeCount = comment.LikeCount - 1
		}
		if newLikeCount < 0 {
			newLikeCount = 0
		}
		if err := DB.Model(&CommentItems{}).Where("id = ?", comment.ID).Update("like_count", newLikeCount).Error; err != nil {
			return err
		}
	}

	return nil
}

func QueryVideosByUserID(userID string, page, pageSize int64) (*[]common.LikeVideoDTO, error) {
	var videos []VideoItems

	tx := DB.Debug().Model(&VideoItems{}).
		Joins("INNER JOIN likes ON likes.likeable_id = videos.id").
		Where("likes.user_id = ? AND likes.likeable_type = ?", userID, "video").
		Order("likes.created_at DESC").
		Limit(int(pageSize)).
		Offset(int(pageSize * (page - 1))).
		Find(&videos)

	if err := tx.Error; err != nil {
		return nil, err
	}

	likeVideoDTOs := make([]common.LikeVideoDTO, 0, len(videos))
	for _, video := range videos {
		var deleteAtStr string
		if video.DeletedAt.Valid {
			deleteAtStr = video.DeletedAt.Time.Format("2006-01-02 15:04:05")
		}

		likeVideoDTOs = append(likeVideoDTOs, common.LikeVideoDTO{
			ID:           video.ID,
			UserID:       video.UserID,
			VideoURL:     video.VideoURL,
			CoverURL:     video.CoverURL,
			Title:        video.Title,
			Description:  video.Description,
			VisitCount:   video.VisitCount,
			LikeCount:    video.LikeCount,
			CommentCount: video.CommentCount,
			CreatedAt:    video.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    video.UpdatedAt.Format("2006-01-02 15:04:05"),
			DeletedAt:    deleteAtStr,
		})
	}

	return &likeVideoDTOs, nil
}

func QueryCommentsByCommentID(commentID string, page, pageSize int64) ([]CommentItems, error) {
	var comments []CommentItems
	tx := DB.Where("parent_id = ?", commentID).Find(&comments)
	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func QueryCommentsByVideoID(videoID string, page, pageSize int64) ([]CommentItems, error) {
	var comments []CommentItems
	tx := DB.Where("video_id = ?", videoID).Find(&comments)
	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func GetVideosByIDLike(videoID string) (VideoItems, error) {
	var video VideoItems
	if err := DB.Where("id = ?", videoID).First(&video).Error; err != nil {
		return VideoItems{}, err
	}
	return video, nil
}

func GetCommentsByIDLike(commentID string) (CommentItems, error) {
	var comment CommentItems
	if err := DB.Where("id = ?", commentID).First(&comment).Error; err != nil {
		return CommentItems{}, err
	}
	return comment, nil
}

func GetVideosByCommentID(commentID string) (string, error) {
	var comment CommentItems
	if err := DB.Where("id = ?", commentID).First(&comment).Error; err != nil {
		return "", err
	}
	if err := DB.Model(&CommentItems{}).Where("id = ?", commentID).Update("child_count", comment.ChildCount+1).Error; err != nil {
		return "", err
	}
	return comment.VideoId, nil
}

func DeleteCommentByCommentID(commentID string) error {
	comment, err := GetCommentsByIDLike(commentID)
	if err != nil {
		return err
	}
	parentID := comment.ParentId
	if parentID != "" {
		// 是二级评论，父评论的子评论数减一
		var parentComment CommentItems
		if err := DB.Where("id = ?", parentID).First(&parentComment).Error; err != nil {
			return err
		}
		if err := DB.Model(&CommentItems{}).Where("id = ?", parentID).Update("child_count", parentComment.ChildCount-1).Error; err != nil {
			return err
		}
	}
	return DB.Where("id = ?", commentID).Delete(&CommentItems{}).Error
}

func DeleteCommentByVideoID(videoID string) error {
	video, err := GetVideosByIDLike(videoID)
	if err != nil {
		return err
	}
	if video.CommentCount > 0 {
		if err := DB.Model(&VideoItems{}).Where("id = ?", videoID).Update("comment_count", 0).Error; err != nil {
			return err
		}
	}
	return DB.Where("video_id = ?", videoID).Delete(&CommentItems{}).Error
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
