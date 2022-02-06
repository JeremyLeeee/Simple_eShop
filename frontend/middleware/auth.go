package middleware

import (
	"github.com/kataras/iris/v12"
)

func AuthConProduct(ctx iris.Context) {
	uid := ctx.GetCookie("uid")
	if uid == "" {
		ctx.Application().Logger().Debug("please login first")
		ctx.Redirect("/user/login")
		return
	}
	ctx.Application().Logger().Debug("already login")
	ctx.Next()
}
