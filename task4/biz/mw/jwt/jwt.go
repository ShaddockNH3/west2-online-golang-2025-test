package jwt

import (
	"context"
	"fmt"
	"time"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/dal/db"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/common"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/model/user"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/configs/constants"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/errno"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v4"
	hertzjwt "github.com/hertz-contrib/jwt"
)

var JwtMiddleware *hertzjwt.HertzJWTMiddleware

func InitJwt() {
	var err error
	JwtMiddleware, err = hertzjwt.New(&hertzjwt.HertzJWTMiddleware{
		Key:           []byte(constants.JwtSecretKey),
		Timeout:       constants.AccessTokenTimeout,
		TokenLookup:   "header: Access-Token, query: token, form: token",
		TokenHeadName: "",
		IdentityKey:   constants.JwtIdentityKey,

		// 登录验证逻辑
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginReq user.LoginUserRequest
			if err := c.BindAndValidate(&loginReq); err != nil {
				return nil, errno.ParamErr.WithMessage(err.Error())
			}
			dbUser, err := db.QueryUserByUsername(loginReq.Username)
			if err != nil {
				return nil, errno.UserNotExistErr
			}
			if ok := utils.VerifyPassword(loginReq.Password, dbUser.Password); !ok {
				return nil, errno.PasswordIsNotVerified
			}
			return dbUser, nil
		},

		// Payload设置逻辑
		PayloadFunc: func(data interface{}) hertzjwt.MapClaims {
			if v, ok := data.(*db.User); ok {
				return hertzjwt.MapClaims{
					constants.JwtIdentityKey: v.ID,
				}
			}
			return hertzjwt.MapClaims{}
		},

		// 登录成功响应逻辑
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, accessToken string, expire time.Time) {
			hlog.CtxInfof(ctx, "用户登录成功, IP: "+c.ClientIP())
			data, _ := c.Get("JWT_PAYLOAD")
			dbUser := data.(*db.User)

			c.Header("Access-Token", accessToken)

			claims := jwt.MapClaims{
				constants.JwtIdentityKey: dbUser.ID,
				"token_type":             "refresh",
				"exp":                    time.Now().Add(constants.RefreshTokenTimeout).Unix(),
				"iat":                    time.Now().Unix(),
			}
			refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			signedRefreshToken, _ := refreshToken.SignedString([]byte(constants.JwtSecretKey))
			c.Header("Refresh-Token", signedRefreshToken)

			createAtStr := dbUser.CreatedAt.Format("2006-01-02 15:04:05")
			updateAtStr := dbUser.UpdatedAt.Format("2006-01-02 15:04:05")
			deleteAtStr := ""
			if dbUser.DeletedAt.Valid {
				deleteAtStr = dbUser.DeletedAt.Time.Format("2006-01-02 15:04:05")
			}

			c.JSON(consts.StatusOK, user.LoginUserResponse{
				Base: &common.BaseResponse{
					Code: fmt.Sprintf("%d", errno.Success.ErrCode),
					Msg:  errno.Success.ErrMsg,
				},
				Data: &common.UserDataResponse{
					ID:        dbUser.ID,
					Username:  dbUser.Username,
					AvatarURL: dbUser.AvatarUrl,
					CreateAt:  createAtStr,
					UpdateAt:  updateAtStr,
					DeleteAt:  deleteAtStr,
				},
			})
		},

		// 验证token逻辑
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if v, ok := data.(string); ok {
				c.Set(constants.ContextCurrentUserKey, v)
				return true
			}
			return false
		},

		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(consts.StatusOK, user.LoginUserResponse{
				Base: &common.BaseResponse{
					Code: "-1",
					Msg:  message,
				},
			})
		},

		// 自定义错误消息格式化
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			err := errno.ConvertErr(e)
			return err.ErrMsg
		},
	})

	if err != nil {
		panic("JWT Middleware 初始化失败: " + err.Error())
	}
}
