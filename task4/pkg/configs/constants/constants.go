package constants

import "time"

// Jwt
const (
	JwtSecretKey          = "task4-secret-key"
	AccessTokenTimeout    = 2 * time.Hour
	RefreshTokenTimeout   = 7 * 24 * time.Hour
	JwtIdentityKey        = "user_id"
	ContextCurrentUserKey = "current_user_id"
)

const (
	MySQLDefaultDSN = "gorm:gorm@tcp(127.0.0.1:9910)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
)

const (
	UserTableName    = "users"
	VideosTableName  = "videos"
	CommentTableName = "comments"
	FollowsTableName = "follows"
)

const (
	Host             = "http://172.28.172.13:8888"
	DefaultURL       = Host + "/static/"
	DefaultAvatarURL = Host + "/static/avatars/default_avatar.jpg"
)
