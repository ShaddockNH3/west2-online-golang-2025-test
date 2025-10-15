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

	video, err := db.QueryVideosByUserIDLike(model.LikeableID)
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
