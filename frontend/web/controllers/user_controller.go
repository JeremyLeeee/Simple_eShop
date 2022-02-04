package controllers

import (
	"eshop/datamodels"
	"eshop/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	Ctx         iris.Context
	UserService services.IUserService
	// session     *sessions.Session
}

func (u *UserController) GetLogin() mvc.View {
	return mvc.View{
		Name: "user/login.html",
	}
}

func (u *UserController) GetRegister() mvc.View {
	return mvc.View{
		Name: "user/register.html",
	}
}

func (u *UserController) PostRegister() {
	var (
		nickName = u.Ctx.FormValue("nickName")
		userName = u.Ctx.FormValue("userName")
		password = u.Ctx.FormValue("password")
	)
	// ozzo validation

	user := &datamodels.User{
		NickName:     nickName,
		UserName:     userName,
		HashPassword: password,
	}

	_, err := u.UserService.AddUser(user)
	if err != nil {
		u.Ctx.Redirect("user/error")
		return
	}
	u.Ctx.Redirect("user/login")
}
