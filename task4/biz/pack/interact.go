package pack

import (
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/dal/db"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/common"
)

func Likes(models []db.LikeItems) []*common.LikeVideoDTO {
	videos := make([]*common.LikeVideoDTO, 0, len(models))
	for _, m := range models {
		videos = append(videos, Like(m))
	}
	return videos
}

func Like(model db.LikeItems) *common.LikeVideoDTO {
	var deleteAtStr string

	video, err := db.GetVideosByIDLike(model.LikeableID)
	if err != nil {
		return &common.LikeVideoDTO{}
	}

	if model.DeletedAt.Valid {
		deleteAtStr = model.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &common.LikeVideoDTO{
		ID:           model.ID,
		UserID:       model.UserID,
		VideoURL:     video.VideoURL,
		CoverURL:     video.CoverURL,
		Title:        video.Title,
		Description:  video.Description,
		VisitCount:   video.VisitCount,
		LikeCount:    video.LikeCount,
		CommentCount: video.CommentCount,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		DeletedAt:    deleteAtStr,
	}
}

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

	return &common.CommentItems{
		ID:         model.ID,
		UserID:     model.UserId,
		VideoID:    model.VideoId,
		ParentID:   model.ParentId,
		LikeCount:  model.LikeCount,
		ChildCount: model.ChildCount,
		Content:    model.Content,
		CreateAt:   model.CreateAt,
		UpdateAt:   model.UpdateAt,
		DeleteAt:   deleteAtStr,
	}
}
