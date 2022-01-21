package controllers

import (
	"product/servies"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type OrderController struct {
	Ctx          iris.Context
	OrderService servies.IOrderService
}

func (o *OrderController) Get() mvc.View {
	orders, err := o.OrderService.GetAllOrderInfo()
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
		o.Ctx.Application().Logger().Debug("failed to get order info")
	}

	return mvc.View{
		Name: "order/view.html",
		Data: iris.Map{
			"order": orders,
		},
	}
}
