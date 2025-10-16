package pack

import (
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/dal/db"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/common"
)

func Comments(models []db.CommentItems) []*common.CommentItems {
	comments := make([]*common.CommentItems, 0, len(models))
	for _, m := range models {
		comments = append(comments, Comment(m))
	}
	return comments
}

func Comment(model db.CommentItems) *common.CommentItems {
	var deleteAtStr string

	if model.DeletedAt.Valid {
		deleteAtStr = model.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	CreateAt := model.CreateAt.Format("2006-01-02 15:04:05")
	UpdateAt := model.UpdateAt.Format("2006-01-02 15:04:05")

	return &common.CommentItems{
		ID:         model.ID,
		UserID:     model.UserId,
		VideoID:    model.VideoId,
		ParentID:   model.ParentId,
		LikeCount:  model.LikeCount,
		ChildCount: model.ChildCount,
		Content:    model.Content,
		CreateAt:   CreateAt,
		UpdateAt:   UpdateAt,
		DeleteAt:   deleteAtStr,
	}
}
