package db

import (
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/constants"
	"gorm.io/gorm"
)

type VideoItems struct {
	ID           string         `gorm:"primaryKey;type:varchar(100)"`
	UserID       string         // user_id
	VideoURL     string         // video_url
	CoverURL     string         // cover_url
	Title        string         // title
	Description  string         // description
	VisitCount   int64          // visit_count
	LikeCount    int64          // like_count
	CommentCount int64          // comment_count
	CreatedAt    string         // create_at
	UpdatedAt    string         // update_at
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (VideoItems) TableName() string {
	return constants.VideosTableName
}

func CreateVideo(video *VideoItems) error {
	return DB.Create(video).Error
}

func QueryVideoByTitle(title string) (*VideoItems, error) {
	var video VideoItems
	if err := DB.Where("title = ?", title).First(&video).Error; err != nil {
		return nil, err
	}
	return &video, nil
}

func QueryVideosByID(userID string, page, pageSize int64) ([]VideoItems, int64, error) {
	var total int64
	if err := DB.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var videos []VideoItems
	if err := DB.Where("user_id = ?", userID).Find(&videos).Error; err != nil {
		return nil, 0, err
	}

	if err := DB.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&videos).Error; err != nil {
		return nil, 0, err
	}

	return videos, total, nil
}

func QueryVideosByKeyword(keyword string, page, pageSize int64, from_date, to_date *int64, username *string) ([]VideoItems, int64, error) {
	if keyword != "" {
		DB = DB.Where(DB.Or("name like ?", "%"+keyword+"%").
			Or("introduce like ?", "%"+keyword+"%"))
	}

	var total int64
	if err := DB.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var videos []VideoItems

	// 页数
	if err := DB.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&videos).Error; err != nil {
		return nil, 0, err
	}

	if from_date != nil {
		if err := DB.Where("created_at >= ?", *from_date).Error; err != nil {
			return nil, 0, err
		}
	}

	if to_date != nil {
		if err := DB.Where("created_at <= ?", *to_date).Error; err != nil {
			return nil, 0, err
		}
	}

	if username != nil {
		user, err := QueryUserByUsername(*username)
		if err != nil {
			return nil, 0, err
		}
		if user.ID != "" {
			if err := DB.Where("user_id = ?", user.ID).Find(&videos).Error; err != nil {
				return nil, 0, err
			}
		} else {
			return nil, 0, nil
		}
	}

	return videos, total, nil
}
