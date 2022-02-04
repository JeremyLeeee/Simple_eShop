package controllers

import (
	"eshop/common"
	"eshop/datamodels"
	"eshop/services"
	"log"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type UserController struct {
	Ctx         iris.Context
	UserService services.IUserService
	Session     *sessions.Session
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
	u.Ctx.Redirect("login")
}

func (u *UserController) PostLogin() mvc.Response {
	var (
		userName = u.Ctx.FormValue("userName")
		password = u.Ctx.FormValue("password")
	)
	log.Println(userName)
	log.Println(password)
	user, isOk := u.UserService.IsPwdSuccess(userName, password)
	log.Print(isOk)
	if !isOk {
		return mvc.Response{
			Path: "login",
		}
	}
	idString := strconv.FormatInt(user.ID, 10)
	common.GlobalCookie(u.Ctx, "uid", idString)
	u.Session.Set("uid", idString)

	return mvc.Response{
		Path: "/product/",
	}
}
