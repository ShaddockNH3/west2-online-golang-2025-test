package interact_service

import (
	"context"
	"errors"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/dal/db"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/common"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/interact"
	"github.com/google/uuid"
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

	if req.CommentID != nil && req.VideoID != nil {
		return errors.New("comment_id and video_id cannot both be set")
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

func (s *InteractService) ListLike(userID string, req *interact.ListLikeRequest) (*[]common.LikeVideoDTO, error) {
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

	likes, err := db.QueryVideosByUserID(userID, currentPageNum, currentPageSize)
	if err != nil {
		return nil, err
	}

	return likes, nil
}

func (s *InteractService) PublishComment(userID string, req *interact.PublishCommentRequest) error {
	var err error
	var currVideoID, currParentID string

	if req.VideoID == nil && req.CommentID == nil {
		return errors.New("video_id and comment_id cannot both be nil")
	}

	if req.VideoID != nil && req.CommentID != nil {
		return errors.New("video_id and comment_id cannot both be set")
	}

	// 如果是直接评论的视频，currParentID为空
	if req.VideoID != nil && *req.VideoID != "" {
		currVideoID = *req.VideoID
		currParentID = ""
	} else if req.CommentID != nil && *req.CommentID != "" { // 回复评论
		currVideoID, err = db.GetVideosByCommentID(*req.CommentID) // 这里直接包含了让子评论+1的逻辑
		if err != nil {
			return err
		}
		currParentID = *req.CommentID
	}

	newComment := &db.CommentItems{
		ID:         uuid.NewString(),
		UserId:     userID,
		VideoId:    currVideoID,
		ParentId:   currParentID,
		LikeCount:  0,
		ChildCount: 0,
		Content:    *req.Content,
	}

	if err = db.CreateComment(newComment); err != nil {
		return err
	}

	return nil
}

func (s *InteractService) ListComment(req *interact.ListCommentRequest) ([]db.CommentItems, error) {
	var err error
	var currentPageSize, currentPageNum int64

	// 判断ID二选一
	if req.VideoID == nil && req.CommentID == nil {
		return nil, errors.New("video_id and comment_id cannot both be nil")
	}

	if req.VideoID != nil && req.CommentID != nil {
		return nil, errors.New("video_id and comment_id cannot both be set")
	}

	// 设置页面参数
	if req.PageNum == nil || *req.PageNum <= 0 {
		currentPageNum = 1
	} else {
		currentPageNum = *req.PageNum
	}
	if req.PageSize == nil || *req.PageSize <= 0 {
		currentPageSize = 10
	} else {
		currentPageSize = *req.PageSize
	}

	// 设置ID
	var id string
	var comments []db.CommentItems

	if req.VideoID != nil {
		id = *req.VideoID
		comments, err = db.QueryCommentsByVideoID(id, currentPageSize, currentPageNum)
	} else if req.CommentID != nil && *req.CommentID != "" {
		id = *req.CommentID
		comments, err = db.QueryCommentsByCommentID(id, currentPageSize, currentPageNum)
	}

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *InteractService) DeleteComment(req *interact.DeleteCommentRequest) error {
	var err error

	if req.CommentID == nil || req.VideoID == nil {
		return errors.New("comment_id and video_id cannot be empty")
	}

	if *req.CommentID == "" || *req.VideoID == "" {
		return errors.New("comment_id and video_id cannot be empty")
	}

	var curr_id string
	if *req.CommentID != "" {
		curr_id = *req.CommentID
		err = db.DeleteCommentByCommentID(curr_id)
	} else {
		curr_id = *req.VideoID
		err = db.DeleteCommentByVideoID(curr_id)
	}

	if err != nil {
		return err
	}

	return nil
}
