package modules

import (
	"eshop/common"
	"eshop/controllers"
	"eshop/repositories"
	"eshop/servies"
	"log"
)

type OrderModule struct {
	name string
	repo repositories.IOrder
}

func NewOrderModule() IModule {
	db, err := common.GetNewGormDB()
	if err != nil {
		log.Print(err)
	}
	repo := repositories.NewOrderManager("order", db)
	return &OrderModule{name: "order", repo: repo}
}

func (m *OrderModule) GetModuleName() string {
	return m.name
}

func (m *OrderModule) GetController() interface{} {
	return &controllers.OrderController{}
}

func (m *OrderModule) GetRepository() interface{} {
	return m.repo
}

func (m *OrderModule) GetService() interface{} {
	return servies.NewOrderService(m.repo)
}
