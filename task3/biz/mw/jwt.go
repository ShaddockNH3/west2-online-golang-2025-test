// 文件路径: biz/mw/jwt.go
package mw

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/dal/mysql"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/model"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/model/task3"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task3/biz/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	UserIDKey     = "UserID"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "mygo_task3_zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour * 24,
		MaxRefresh:    time.Hour * 24,
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",

		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, &task3.LoginResponse{
				Code:  task3.Code_Success,
				Msg:   "login success",
				Token: &token,
			})
		},

		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var req task3.LoginRequest
			if err := c.BindAndValidate(&req); err != nil {
				return nil, err
			}

			user, err := mysql.QueryUserByName(req.Name)
			if err != nil {
				return nil, err
			}
			if user == nil {
				return nil, errors.New("user not found")
			}

			if err := utils.CheckPassword(user.Password, req.Password); err != nil {
				return nil, errors.New("wrong password")
			}

			return user, nil
		},

		IdentityKey: UserIDKey,

		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return int64(claims[UserIDKey].(float64))
		},

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					UserIDKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},

		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusUnauthorized, &task3.LoginResponse{
				Code: task3.Code(task3.Code_ParamInvalid), // 或者可以定义一个新的 AuthFailed 的 Code
				Msg:  message,
			})
		},

		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
	})
	if err != nil {
		panic(err)
	}
}
