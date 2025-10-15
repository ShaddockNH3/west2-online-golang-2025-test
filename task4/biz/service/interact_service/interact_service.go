package interact_service

import (
	"context"
	"errors"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/dal/db"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/interact"
)

type InteractService struct {
	ctx context.Context
}

func NewInteractService(ctx context.Context) *InteractService {
	return &InteractService{ctx: ctx}
}

func (s *InteractService) ActionLike(userID string, req *interact.ActionLikeRequest) error {
	var err error

	if req.CommentID == nil && req.VideoID == nil {
		return errors.New("comment_id and video_id cannot both be nil")
	}

	if *req.CommentID == "" && *req.VideoID == "" {
		return errors.New("comment_id and video_id cannot both be empty")
	}

	var likeableID string
	var likeType int64

	if req.ActionType == nil {
		likeType = 1
	} else {
		likeType = int64(*req.ActionType)
	}

	if req.CommentID != nil && *req.CommentID != "" {
		likeableID = *req.CommentID
		err = db.UpdateLike(likeableID, "comment", userID, likeType)

	} else if req.VideoID != nil && *req.VideoID != "" {
		likeableID = *req.VideoID
		err = db.UpdateLike(likeableID, "video", userID, likeType)
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *InteractService) ListLike(userID string, req *interact.ListLikeRequest) ([]db.LikeItems, error) {
	var err error
	var currentPageSize, currentPageNum int64

	if req.PageNum == nil {
		currentPageNum = 1
	} else {
		currentPageNum = *req.PageNum
	}

	if req.PageSize == nil {
		currentPageSize = 10
	} else {
		currentPageSize = *req.PageSize
	}

	likes, err := db.QueryVideosByUserID(userID, currentPageSize, currentPageNum)
	if err != nil {
		return nil, err
	}

	return likes, nil
}
