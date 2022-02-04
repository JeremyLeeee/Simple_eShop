package main

import (
	"context"
	"eshop/common"
	"eshop/frontend/web/controllers"
	"eshop/repositories"
	"eshop/services"
	"log"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func main() {
	// create iris instance
	app := iris.New()
	// set log level
	app.Logger().SetLevel("debug")
	// register templates
	template := iris.HTML("./web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	// set template target
	app.HandleDir("/public", "./web/public")
	// jump to error page while error occur
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "error on page!"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	db, err := common.NewMysqlConn()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sess := sessions.New(sessions.Config{
		Cookie:  "helloworld",
		Expires: 60 * time.Minute,
	})

	// register controller
	userRepository := repositories.NewUserRepository("user", db)
	userService := services.NewUserService(userRepository)
	userParty := app.Party("/user")
	user := mvc.New(userParty)
	user.Register(ctx, userService, sess.Start)
	user.Handle(new(controllers.UserController))

	// start service
	app.Run(
		iris.Addr("localhost:8082"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}
