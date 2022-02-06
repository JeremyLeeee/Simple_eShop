package controllers

import (
	"eshop/datamodels"
	"eshop/services"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.IProductService
	OrderService   services.IOrderService
	Session        *sessions.Session
}

func (p *ProductController) GetDetail() mvc.View {
	product, err := p.ProductService.GetProductById(5)
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

func (p *ProductController) GetOrder() mvc.View {
	productIDString := p.Ctx.URLParam("productID")
	userID := p.Ctx.GetCookie("uid")
	productID, err := strconv.Atoi(productIDString)
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	product, err := p.ProductService.GetProductById(int64(productID))
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}

	var orderID int64
	if product.ProductNum > 0 {
		// will be optimized later
		product.ProductNum -= 1
		err := p.ProductService.UpdateProduct(product)
		if err != nil {
			p.Ctx.Application().Logger().Error(err)
		}
		// create order
		userID, err := strconv.Atoi(userID)
		if err != nil {
			p.Ctx.Application().Logger().Error(err)
		}
		order := &datamodels.Order{
			UserId:      int64(userID),
			ProductId:   int64(productID),
			OrderStatus: datamodels.OrderSuccess,
		}
		// insert order
		orderID, err = p.OrderService.InsertOrder(order)
		if err != nil {
			p.Ctx.Application().Logger().Error(err)
		}
		return mvc.View{
			Layout: "shared/productLayout.html",
			Name:   "product/result.html",
			Data: iris.Map{
				"showMessage": "Ordered successfully",
				"orderID":     orderID,
			},
		}
	} else {
		return mvc.View{
			Layout: "shared/productLayout.html",
			Name:   "product/result.html",
			Data: iris.Map{
				"showMessage": "failed to buy",
				"orderID":     orderID,
			},
		}
	}

}
