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
	MySQLDefaultDSN = "west2:west2_password@tcp(127.0.0.1:3306)/west2_test?charset=utf8mb4&parseTime=True&loc=Local"
)

const (
	UserTableName    = "users"
	VideosTableName  = "videos"
	CommentTableName = "comments"
	FollowsTableName = "follows"
)

const (
	DefaultAvatarURL = "https://ShaddockNH3.github.com/static/default_avatar.png"
)
