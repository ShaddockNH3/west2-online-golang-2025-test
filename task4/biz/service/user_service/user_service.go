package user_service

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/dal/db"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/user"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/mw/jwt"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/constants"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/errno"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/utils"
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
	if err != nil && err.Error() != "record not found" {
		return err
	}
	if user != nil {
		return errno.UserAlreadyExistErr
	}

	// password加密
	passwordHash, err := utils.Crypt(req.Password)
	if err != nil {
		return err
	}

	newUser := &db.User{
		ID:        uuid.NewString(),
		Username:  req.Username,
		Password:  passwordHash,
		AvatarUrl: constants.DefaultAvatarURL,
	}

	if err = db.CreateUser(newUser); err != nil {
		return err
	}

	return nil
}

func (s *UserService) LoginUser(req *user.LoginUserRequest) (*db.User, string, string, error) {
	var err error

	user, err := db.QueryUserByUsername(req.Username)
	if err != nil {
		return nil, "", "", err
	}
	if user == nil {
		return nil, "", "", errno.UserNotExistErr
	}

	// password验证
	currentUser, err := db.QueryUserByUsername(req.Username)
	if err != nil {
		return nil, "", "", err
	}
	if currentUser == nil {
		return nil, "", "", errno.UserNotExistErr
	}

	if !utils.VerifyPassword(req.Password, currentUser.Password) {
		return nil, "", "", errno.PasswordIsNotVerified
	}

	accessToken, _, err := jwt.AccessTokenJwtMiddleware.TokenGenerator(user)
	if err != nil {
		return nil, "", "", err
	}
	refreshToken, _, err := jwt.RefreshTokenJwtMiddleware.TokenGenerator(user)
	if err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

func (s *UserService) InfoUser(userID string, req *user.InfoUserRequest) (*db.User, error) {
	var err error

	user, err := db.QueryUserByUserId(userID)

	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errno.UserNotExistErr
	}

	return user, nil
}

func (s *UserService) AvatarUploadUser(user_id string, fileHeader *multipart.FileHeader, savePath string, req *user.AvatarUploadUserRequest) (*db.User, error) {
	var err error

	newImageID := uuid.NewString()
	AvatarURL := constants.DefaultURL + "avatars/" + newImageID

	newImage := &db.Image{
		ID:               newImageID,
		UserID:           user_id,
		URL:              AvatarURL,
		OriginalFilename: fileHeader.Filename,
		Filepath:         savePath,
		Filesize:         fileHeader.Size,
		MimeType:         fileHeader.Header.Get("Content-Type"),
	}

	if err = db.CreateImage(newImage); err != nil {
		return nil, err
	}

	if err := db.UploadAvatar(user_id, AvatarURL); err != nil {
		return nil, err
	}

	updatedUser, err := db.QueryUserByUserId(user_id)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *UserService) SearchImage(filename string, req *user.SearchImageRequest) (string, error) {
	var err error

	fileURL, err := db.QueryImageByFilename(filename)
	if err != nil {
		return "", err
	}

	return fileURL, nil
}
