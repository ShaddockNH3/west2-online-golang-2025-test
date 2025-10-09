package service

import (
	"context"
	"errors"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/dal/mysql"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/model"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/model/task3"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/utils"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (s *UserService) UpdateUser(currentUserID int64, req *task3.UpdateUserRequest) error {
	var err error

	if currentUserID != req.UserID {
		return errors.New("permission denied: you can only update your own information")
	}

	userToUpdate := &model.User{
		Name:      req.Name,
		Introduce: req.Introduce,
	}

	if err = mysql.UpdateUser(userToUpdate); err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(currentUserID int64, req *task3.DeleteUserRequest) error {
	var err error

	if currentUserID != req.UserID {
		return errors.New("permission denied: you can only delete your own profile")
	}

	if err = mysql.DeleteUser(req.UserID); err != nil {
		return err
	}

	return nil
}

func (s *UserService) QueryUser(req *task3.QueryUserRequest) ([]*model.User, int64, error) {
	var err error

	users, total, err := mysql.QueryUser(req.Keyword, req.Page, req.PageSize)

	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (s *UserService) CreateUser(req *task3.CreateUserRequest) error {
	var err error

	users, err := mysql.QueryUserByName(req.Name)
	if err != nil {
		return err
	}

	if users != nil {
		return err
	}

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	newUser := &model.User{
		Name:      req.Name,
		Introduce: req.Introduce,
		Password:  password,
	}

	if err = mysql.CreateUser(newUser); err != nil {
		return err
	}
	return nil
}
