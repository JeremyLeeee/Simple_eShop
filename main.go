package main

import (
	"context"
	"eshop/modules"
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	// create iris instance
	app := iris.New()
	// set log level
	app.Logger().SetLevel("debug")
	// register templates
	template := iris.HTML("./web/backend/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	// set template target
	app.HandleDir("/assets", "./web/backend/assets")
	// jump to error page while error occur
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "error on page!"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// register controller
	order := modules.NewOrderModule()
	RegisterMVC(order, app, ctx)

	product := modules.NewProductModule()
	RegisterMVC(product, app, ctx)

	// start service
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
	)

}

func RegisterMVC(module modules.IModule, app *iris.Application, ctx context.Context) {
	log.Println("Register " + module.GetModuleName())

	service := module.GetService()
	controller := module.GetController()

	party := app.Party("/" + module.GetModuleName())

	MVC := mvc.New(party)
	MVC.Register(ctx, service)
	MVC.Handle(controller)
}
