package db

import (
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/mw/redis"
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

func UpdateVisitCount(videoID string, count int64) error {
	return DB.Model(&VideoItems{}).Where("id = ?", videoID).Update("visit_count", count).Error
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
	var err error
	tx := DB.Model(&VideoItems{})

	if keyword != "" {
		tx = tx.Where(tx.Or("name like ?", "%"+keyword+"%").
			Or("introduce like ?", "%"+keyword+"%"))
	}

	if from_date != nil {
		if err = tx.Where("created_at >= ?", *from_date).Error; err != nil {
			return nil, 0, err
		}
	}

	if to_date != nil {
		if err = tx.Where("created_at <= ?", *to_date).Error; err != nil {
			return nil, 0, err
		}
	}

	if username != nil {
		user, err := QueryUserByUsername(*username)
		if err != nil {
			return nil, 0, err
		}
		if user.ID == "" {
			return []VideoItems{}, 0, nil
		}
		tx = tx.Where("user_id = ?", user.ID)
	}

	var total int64
	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var videos []VideoItems

	// 页数
	queryChain := tx
	if err := queryChain.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&videos).Error; err != nil {
		return nil, 0, err
	}

	if len(videos) == 0 {
		return videos, total, nil
	}

	// 批量更新浏览量
	videoIDs := make([]string, len(videos))
	for i, v := range videos {
		videoIDs[i] = v.ID
	}
	if err := DB.Model(&VideoItems{}).Where("id IN (?)", videoIDs).Update("visit_count", gorm.Expr("visit_count + 1")).Error; err != nil {
		// return nil, 0, err
	}

	// 更新热门视频
	go func() {
		for _, video := range videos {
			newVisitCount := video.VisitCount + 1
			if redis.CheckRdbPopular(video.ID) { // 已存在，更新
				redis.AddRdbPopular(video.ID, newVisitCount)
			}
		}
	}()

	return videos, total, nil
}

func PopularVideos(page, pageSize int64) ([]VideoItems, error) {
	start := (page - 1) * pageSize
	end := start + pageSize - 1

	videoIDs, err := redis.GetHotVideoIDs(constants.PopularVideosSuffix, start, end)
	if err != nil {
		// 假设 Redis 是可靠的，后续可以进行数据库降级处理
		return nil, err
	}

	if len(videoIDs) == 0 {
		return []VideoItems{}, nil
	}

	// 根据 videoIDs 顺序查询视频信息
	var videos []VideoItems
	if err := DB.Order(gorm.Expr("FIELD(id, ?)", videoIDs)).Where("id IN (?)", videoIDs).Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}
