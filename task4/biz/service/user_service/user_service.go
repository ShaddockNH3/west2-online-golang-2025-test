package user_service

import (
	"context"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/dal/db"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/user"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/configs/constants"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/errno"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/utils"
	"github.com/google/uuid"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (s *UserService) RegisterUser(req *user.RegisterUserRequest) error {
	var err error

	user, err := db.QueryUserByUsername(req.Username)
	if err != nil {
		return err
	}
	if user != nil {
		return errno.UserAlreadyExistErr
	}

	//password加密
	passwordHash, err := utils.Crypt(req.Password)
	if err != nil {
		return err
	}

	newUser := &db.User{
		ID:        uuid.NewString(),
		Username:  req.Username,
		Password:  passwordHash,
		AvatarUrl: constants.DefaultAvatarURL,
		// 默认头像还需要自己再编写
	}

	if err = db.CreateUser(newUser); err != nil {
		return err
	}

	return nil
}

func (s *UserService) InfoUser(req *user.InfoUserRequest) (db.User, error) {
	var err error

	user, err := db.QueryUserByUserId(*req.UserID)

	if err != nil {
		return db.User{}, err
	}
	if user == nil {
		return db.User{}, errno.UserNotExistErr
	}

	return *user, nil
}
