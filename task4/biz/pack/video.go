package pack

import (
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/dal/db"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/common"
)

func Videos(models []db.VideoItems) []*common.VideoItems {
	videos := make([]*common.VideoItems, 0, len(models))
	for _, m := range models {
		videos = append(videos, Video(m))
	}
	return videos
}

func Video(model db.VideoItems) *common.VideoItems {
	var deleteAtStr string
	if model.DeletedAt.Valid {
		deleteAtStr = model.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	CreateAt := model.CreatedAt.Format("2006-01-02 15:04:05")
	UpdateAt := model.UpdatedAt.Format("2006-01-02 15:04:05")

	return &common.VideoItems{
		ID:           model.ID,
		UserID:       model.UserID,
		VideoURL:     model.VideoURL,
		CoverURL:     model.CoverURL,
		Title:        model.Title,
		Description:  model.Description,
		VisitCount:   model.VisitCount,
		LikeCount:    model.LikeCount,
		CommentCount: model.CommentCount,
		CreateAt:     CreateAt,
		UpdateAt:     UpdateAt,
		DeleteAt:     deleteAtStr,
	}
}
