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

func (s *VideoService) ListVideos(req *video.ListVideoRequest) ([]db.VideoItems, int64, error) {
	videos, total, err := db.QueryVideosByID(req.UserID, req.PageNum, req.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return videos, total, nil
}

func (s *VideoService) SearchVideos(req *video.SearchVideoRequest) ([]db.VideoItems, int64, error) {
	videos, total, err := db.QueryVideosByKeyword(req.Keyword, req.PageNum, req.PageSize, req.FromDate, req.ToDate, req.Username)
	if err != nil {
		return nil, 0, err
	}
	return videos, total, nil
}

func (s *VideoService) PopularVideos(req *video.PopularVideoRequest) ([]db.VideoItems, error) {
	var currentPage int64
	var pageSize int64

	if req.PageNum == nil || *req.PageNum <= 0 {
		currentPage = 1
	}

	if req.PageSize == nil || *req.PageSize <= 0 {
		pageSize = 10
	}

	videos, err := db.PopularVideos(currentPage, pageSize)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
