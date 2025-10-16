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
	// MySQLDefaultDSN = "gorm:gorm@tcp(127.0.0.1:9910)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	MySQLDefaultDSN = "gorm:gorm@tcp(mysql:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// RedisAddr     = "127.0.0.1:9911"
	RedisAddr     = "redis:6379"
	RedisPassword = "shenmidazhi"
)

const (
	UserTableName     = "users"
	VideosTableName   = "videos"
	LikesTableName    = "likes"
	CommentsTableName = "comments"
	FollowsTableName  = "follows"
)

const (
	// Host             = "http://172.28.172.13:8888"
	Host             = "http://localhost:8080" // 根据实际情况修改为服务器地址和端口
	DefaultURL       = Host + "/static/"
	DefaultAvatarURL = Host + "/static/avatars/default_avatar.jpg"
)

const (
	PopularVideosSuffix = ":popular_videos"
)
