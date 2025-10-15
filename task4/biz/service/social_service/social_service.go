package social_service

import (
	"context"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/dal/db"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/common"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/social"
)

type SocialService struct {
	ctx context.Context
}

func NewSocialService(ctx context.Context) *SocialService {
	return &SocialService{ctx: ctx}
}

func (s *SocialService) ActionRelation(userID string, req *social.ActionRelationRequest) error {
	var err error

	err = db.UpdateRelation(userID, *req.ToUserID, int64(*req.ActionType))
	if err != nil {
		return err
	}

	return nil
}

func (s *SocialService) GetFollowList(req *social.ListFollowingRequest) ([]*common.SocialDTO, error) {
	var err error
	var currentPageNum, currentPageSize int64

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

	users, err := db.QueryFollowingList(req.UserID, currentPageNum, currentPageSize)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *SocialService) GetFollowerList(req *social.ListFollowerRequest) ([]*common.SocialDTO, error) {
	var err error
	var currentPageNum, currentPageSize int64
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

	users, err := db.QueryFollowerList(req.UserID, currentPageNum, currentPageSize)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *SocialService) GetFriendList(userID string, req *social.ListFriendsRequest) ([]*common.SocialDTO, error) {
	var err error
	var currentPageNum, currentPageSize int64
	// 设置页面参数
	if req.PageNum == nil || *req.PageNum <= 0 {
		currentPageNum = 1
	}else {
		currentPageNum = *req.PageNum
	}
	if req.PageSize == nil || *req.PageSize <= 0 {
		currentPageSize = 10
	}else {
		currentPageSize = *req.PageSize
	}

	users, err := db.QueryFriendList(userID, currentPageNum, currentPageSize)
	if err != nil {
		return nil, err
	}

	return users, nil
}
