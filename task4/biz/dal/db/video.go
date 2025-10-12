package db

import (
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/configs/constants"
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