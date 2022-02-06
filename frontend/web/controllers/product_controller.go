package controllers

import (
	"eshop/services"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.IProductService
	Session        *sessions.Session
}

func (p *ProductController) GetDetail() mvc.View {
	id, err := strconv.ParseInt(p.Ctx.URLParam("productID"), 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	product, err := p.ProductService.GetProductById(id)
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": product,
		},
	}
}
