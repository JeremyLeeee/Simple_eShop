package main

import (
	"context"
	"fmt"
	"product/common"
	"product/controllers"
	"product/repositories"
	"product/servies"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
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
	app.HandleDir("/assets", "./web/assets")
	// jump to error page while error occur
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "error on page!"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	db, err := common.NewMysqlConn()

	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// register controller
	productRepo := repositories.NewProductManager("product", db)
	productService := servies.NewProductService(productRepo)
	productParty := app.Party("/product")
	productMVC := mvc.New(productParty)
	productMVC.Register(ctx, productService)
	productMVC.Handle(new(controllers.ProductController))

	orderRepo := repositories.NewOrderManager("order", db)
	orderService := servies.NewOrderService(orderRepo)
	orderParty := app.Party("/order")
	orderMVC := mvc.New(orderParty)
	orderMVC.Register(ctx, orderService)
	orderMVC.Handle(new(controllers.OrderController))

	// start service
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
	)

}
