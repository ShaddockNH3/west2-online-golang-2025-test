package user_service

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/dal/db"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/user"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/mw/jwt"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/mw/redis"
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

	// password加密，采用bcrypt，这是一种单向加密
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

	if user.MfaSecret != "" {
		if req.Code == nil {
			return nil, "", "", errno.MFACodeEmptyErr
		}
		// mfa验证
		isValid := totp.Validate(*req.Code, user.MfaSecret)
		if !isValid {
			return nil, "", "", errno.MFAInvalidCodeErr
		}
	}

	// password验证
	if !utils.VerifyPassword(req.Password, user.Password) {
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

func (s *UserService) QrcodeMFAAuth(userID string, req *user.QrcodeMFAAuthRequest) (string, string, error) {
	var err error

	user, err := db.QueryUserByUserId(userID)
	if err != nil {
		return "", "", err
	}
	if user == nil {
		return "", "", errno.UserNotExistErr
	}

	if user.MfaSecret != "" {
		return "", "", errno.MFAAlreadyEnabledErr
	}

	// 生成mfaSecret和二维码URL
	mfaSecret, qrCodeURL, err := utils.GenerateMFAData(user.Username)
	if err != nil {
		return "", "", err
	}

	// 存储mfaSecret到数据库
	// err = db.DB.Model(&db.User{}).Where("id = ?", userID).Update("mfa_secret", mfaSecret).Error
	// if err != nil {
	// 	return "", "", err
	// }

	// 应该是校验完毕了再存入数据库里，这部分内容存入redis里比较好，因为是实时需要的

	redis.SetMFASecretToCache(userID, mfaSecret)

	return qrCodeURL, mfaSecret, nil
}

func (s *UserService) BindMFAAuth(userID string, req *user.BindMFAAuthRequest) error {
	var err error

	if req.Code == nil && req.Secret == nil {
		return errno.MFACodeEmptyErr
	}

	user, err := db.QueryUserByUserId(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errno.UserNotExistErr
	}

	// 从缓存中获取mfaSecret
	mfaSecret, err := redis.GetMFASecretFromCache(userID)
	if err != nil {
		return err
	}

	if req.Code != nil {
		isValid := totp.Validate(*req.Code, mfaSecret)
		if !isValid {
			return errno.MFAInvalidCodeErr
		}
	} else {
		if *req.Secret != mfaSecret {
			return errno.MFAInvalidSecretErr
		}
	}

	// 应该是要实现加密逻辑的，但是得使用双向的加密，而不是使用bcrypt这种单向加密
	// 但是目前先不实现了
	
	// mfaSecretHash, err := utils.Encrypt(mfaSecret)
	// if err != nil {
	// 	return err
	// }

	// 存储mfaSecret到数据库
	err = db.DB.Model(&db.User{}).Where("id = ?", userID).Update("mfa_secret", mfaSecret).Error
	if err != nil {
		return err
	}

	// 从缓存中删除mfaSecret
	redis.DelMFASecretFromCache(userID)

	return nil
}

func (s *UserService) SearchImage(filename string, req *user.SearchImageRequest) (string, error) {
	var err error

	fileURL, err := db.QueryImageByFilename(filename)
	if err != nil {
		return "", err
	}

	return fileURL, nil
}
