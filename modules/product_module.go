package modules

import (
	"eshop/common"
	"eshop/controllers"
	"eshop/repositories"
	"eshop/servies"
	"log"
)

type ProductModule struct {
	name string
	repo repositories.IProduct
}

func NewProductModule() IModule {
	db, err := common.GetNewGormDB()
	if err != nil {
		log.Print(err)
	}
	repo := repositories.NewProductManager("product", db)
	return &ProductModule{name: "product", repo: repo}
}

func (m *ProductModule) GetModuleName() string {
	return m.name
}

func (m *ProductModule) GetController() interface{} {
	return &controllers.ProductController{}
}

func (m *ProductModule) GetRepository() interface{} {
	return m.repo
}

func (m *ProductModule) GetService() interface{} {
	return servies.NewProductService(m.repo)
}
