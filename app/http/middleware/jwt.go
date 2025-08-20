package middleware

import (
	"errors"
	"github.com/goravel/framework/auth"
	contractshttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Jwt() contractshttp.Middleware {
	return func(ctx contractshttp.Context) {

		//获取header中的Authorization 为Bearer token
		token := ctx.Request().Header("Authorization", "")
		//如果token为空
		if token == "" {
			ctx.Request().AbortWithStatusJson(401, contractshttp.Json{
				"message": "请先登录",
			})
			return
		}
		token = token[7:]

		_, err := facades.Auth(ctx).Parse(token)
		//fmt.Println(payload)
		if err != nil {
			ctx.Request().AbortWithStatusJson(401, contractshttp.Json{
				"message": "请先登录",
			})
			return
		}
		is := errors.Is(err, auth.ErrorTokenExpired)
		if is {
			ctx.Request().AbortWithStatusJson(401, contractshttp.Json{
				"message": "登录过期",
			})
			return
		}
		ctx.Request().Next()
	}
}
