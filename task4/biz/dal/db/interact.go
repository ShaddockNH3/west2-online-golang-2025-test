package db

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/common"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/mw/redis"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/constants"
)

type LikeItems struct {
	ID     string `gorm:"primaryKey;type:varchar(100)"`
	UserID string // user_id，指的是点赞用户的

	LikeableID   string `gorm:"index"` // 被点赞对象视频ID或评论ID
	LikeableType string `gorm:"index"` // 被点赞对象的类型 "video" 或 "comment"

	CreatedAt time.Time      // created_at
	UpdatedAt time.Time      // updated_at
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
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (LikeItems) TableName() string {
	return constants.LikesTableName
}

func (CommentItems) TableName() string {
	return constants.CommentsTableName
}

func UpdateLike(userID, likeableType, likeableID string, likeAction int64) error {
	if likeableType != "video" && likeableType != "comment" {
		return errors.New("invalid likeable type")
	}

	if likeAction == 1 { // 点赞
		isNewLike, err := redis.AddLike(userID, likeableType, likeableID)
		if err != nil {
			return err
		}
		// 如果是新点赞，创建数据库记录
		if isNewLike {
			go persistLike(userID, likeableType, likeableID)
		}
	}else if likeAction == 2 { // 取消点赞
		wasLiked, err := redis.RemoveLike(userID, likeableType, likeableID)
		if err != nil {
			return err
		}
		// 如果是成功取消点赞，删除数据库记录
		if wasLiked {
			go persistUnlike(userID, likeableType, likeableID)
		}else{
			return errors.New("not liked yet")
		}
	}
	return nil
}

func persistLike(userID, likeableType, likeableID string) {
	var likeRecord LikeItems
	// Unscoped() 可以在查找时也包括被软删除的记录
	err := DB.Unscoped().Where("user_id = ? AND likeable_id = ?", userID, likeableID).First(&likeRecord).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 记录不存在，创建新的
		newLike := &LikeItems{
			ID:           uuid.New().String(),
			UserID:       userID,
			LikeableID:   likeableID,
			LikeableType: likeableType,
		}
		DB.Create(newLike)
	} else if err == nil && likeRecord.DeletedAt.Valid {
		// 记录存在但被软删除了，恢复它
		DB.Unscoped().Model(&likeRecord).Update("deleted_at", nil)
	}
	// 如果记录存在且未被删除，则什么都不做，因为Redis已经处理过了

	// 更新视频/评论表里的总赞数
	// 这一步是为了数据最终一致性，即使Redis数据丢失也能恢复
	var countFieldToUpdate string
	var modelToUpdate interface{}
	if likeableType == "video" {
		modelToUpdate = &VideoItems{}
		countFieldToUpdate = "like_count"
	} else {
		modelToUpdate = &CommentItems{}
		countFieldToUpdate = "like_count"
	}
	DB.Model(modelToUpdate).Where("id = ?", likeableID).Update(countFieldToUpdate, gorm.Expr(countFieldToUpdate+" + 1"))
}

// persistUnlike - 异步持久化取消点赞记录到MySQL
func persistUnlike(userID, likeableType, likeableID string) {
	// 直接软删除，不用判断是否存在，因为能触发这个函数说明Redis里是存在的
	DB.Where("user_id = ? AND likeable_id = ?", userID, likeableID).Delete(&LikeItems{})

	// 更新视频/评论表里的总赞数
	var countFieldToUpdate string
	var modelToUpdate interface{}
	if likeableType == "video" {
		modelToUpdate = &VideoItems{}
		countFieldToUpdate = "like_count"
	} else {
		modelToUpdate = &CommentItems{}
		countFieldToUpdate = "like_count"
	}
	DB.Model(modelToUpdate).Where("id = ? AND "+countFieldToUpdate+" > 0", likeableID).Update(countFieldToUpdate, gorm.Expr(countFieldToUpdate+" - 1"))
}

func CreateComment(comment *CommentItems) error {
	return DB.Create(comment).Error
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

	err := DB.Where("parent_id = ?", commentID).
		Limit(int(pageSize)).
		Offset(int(pageSize * (page - 1))).
		Find(&comments).Error

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func QueryCommentsByVideoID(videoID string, page, pageSize int64) ([]CommentItems, error) {
	var comments []CommentItems

	err := DB.Where("video_id = ?", videoID).
		Limit(int(pageSize)).
		Offset(int(pageSize * (page - 1))).
		Find(&comments).Error

	if err != nil {
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

func GetUserIdByID(ID string) (string, error) {
	var comment CommentItems
	if err := DB.Where("id = ?", ID).First(&comment).Error; err != nil {
		return "", err
	}
	return comment.UserId, nil
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
