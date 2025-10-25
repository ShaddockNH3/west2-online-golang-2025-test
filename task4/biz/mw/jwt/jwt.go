package jwt

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	hertzjwt "github.com/hertz-contrib/jwt"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/dal/db"
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/constants"
)

var (
	AccessTokenJwtMiddleware  *hertzjwt.HertzJWTMiddleware
	RefreshTokenJwtMiddleware *hertzjwt.HertzJWTMiddleware
)

func InitJwt() {
	var err error
	AccessTokenJwtMiddleware, err = hertzjwt.New(&hertzjwt.HertzJWTMiddleware{
		Key:           []byte(constants.JwtSecretKey),
		Timeout:       constants.AccessTokenTimeout,
		TokenLookup:   "header:Access-Token",
		TokenHeadName: "Bearer",
		IdentityKey:   constants.JwtIdentityKey,

		// Payload设置逻辑
		PayloadFunc: func(data interface{}) hertzjwt.MapClaims {
			if v, ok := data.(*db.User); ok {
				return hertzjwt.MapClaims{
					constants.JwtIdentityKey: v.ID,
				}
			}
			return hertzjwt.MapClaims{}
		},

		// 身份验证逻辑
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := hertzjwt.ExtractClaims(ctx, c)
			if id, ok := claims[constants.JwtIdentityKey]; ok {
				return id
			}
			return nil
		},
	})

	if err != nil {
		panic("JWT Middleware 初始化失败: " + err.Error())
	}

	RefreshTokenJwtMiddleware, err = hertzjwt.New(&hertzjwt.HertzJWTMiddleware{
		Key:           []byte(constants.JwtSecretKey),
		Timeout:       constants.RefreshTokenTimeout,
		TokenLookup:   "header:Refresh-Token",
		TokenHeadName: "Bearer",
		IdentityKey:   constants.JwtIdentityKey,

		// Payload设置逻辑
		PayloadFunc: func(data interface{}) hertzjwt.MapClaims {
			if v, ok := data.(*db.User); ok {
				return hertzjwt.MapClaims{
					constants.JwtIdentityKey: v.ID,
				}
			}
			return hertzjwt.MapClaims{}
		},

		// 身份验证逻辑
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := hertzjwt.ExtractClaims(ctx, c)
			if user_id, ok := claims[constants.JwtIdentityKey]; ok {
				return user_id
			}
			return nil
		},
	})

	if err != nil {
		panic("JWT Middleware 初始化失败: " + err.Error())
	}
}
