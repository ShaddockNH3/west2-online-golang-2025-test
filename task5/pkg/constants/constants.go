package constants

import "time"

// Jwt
const (
	AppName               = "west-online-golang-2025-test-task5"
	JwtSecretKey          = "task5-jwt-secret-key"
	AccessTokenTimeout    = 2 * time.Hour
	RefreshTokenTimeout   = 7 * 24 * time.Hour
	JwtIdentityKey        = "user_id"
	ContextCurrentUserKey = "user_id"
)

// MFA
const (
	MfaSecretKey = "86480212befaa02e32122bd55f35297054b4580025d5880a6247756091dbf6f1" // 后续应该置入环境变量
)

const (
	// MySQLDefaultDSN = "gorm:gorm@tcp(127.0.0.1:9910)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	MySQLDefaultDSN = "gorm:gorm@tcp(mysql:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// RedisAddr = "127.0.0.1:9911"
	RedisAddr     = "redis:6379"
	RedisPassword = "shenmidazhi"
)

const (
	UserTableName     = "users"
	ImageTableName    = "images"
	VideosTableName   = "videos"
	LikesTableName    = "likes"
	CommentsTableName = "comments"
	FollowsTableName  = "follows"
)

const (
	// Host = "http://172.28.172.13:8888"
	Host             = "http://localhost:8080" // 根据实际情况修改为服务器地址和端口
	DefaultURL       = Host + "/static/"
	DefaultAvatarURL = Host + "/static/avatars/default_avatar.jpg"
)

const (
	PopularVideosSuffix   = ":popular_videos"
	MFAAuthSuffix         = ":mfa_auth"
	VideoLikeAuthSuffix   = ":video_like_auth"
	CommentLikeAuthSuffix = ":comment_like_auth"
)
