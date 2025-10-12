package video_service

import (
	"context"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/dal/db"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/video"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/errno"
	"github.com/google/uuid"
)

type VideoService struct {
	ctx context.Context
}

func NewVideoService(ctx context.Context) *VideoService {
	return &VideoService{ctx: ctx}
}

func (s *VideoService) CreateVideo(UserID string, VideoURL string, CoverURL string, Title string, Description string, req *video.PublishVideoRequest) error {
	video, err := db.QueryVideoByTitle(Title)
	if err != nil {
		return err
	}
	if video != nil {
		return errno.VideoAlreadyExistErr
	}

	newVideo := &db.VideoItems{
		ID:           uuid.NewString(),
		UserID:       UserID,
		VideoURL:     VideoURL,
		CoverURL:     CoverURL,
		Title:        Title,
		Description:  Description,
		VisitCount:   0,
		LikeCount:    0,
		CommentCount: 0,
	}

	if err = db.CreateVideo(newVideo); err != nil {
		return err
	}

	return nil
}
