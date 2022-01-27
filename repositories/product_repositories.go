package repositories

import (
	"eshop/common"
	"eshop/datamodels"
	"log"

	"gorm.io/gorm"
)

type IProduct interface {
	// database connection
	Conn() error
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

type ProductManager struct {
	table string
	db    *gorm.DB
}

func NewProductManager(table string, db *gorm.DB) IProduct {
	return &ProductManager{table: table, db: db}
}

func (p *ProductManager) Conn() (err error) {
	if p.db == nil {
		db, err := common.GetNewGormDB()
		if err != nil {
			log.Fatalln("------Error on connection using gorm")
			return err
		}
		p.db = db
	}
	if p.table == "" {
		p.table = "products"
	}
	p.db.AutoMigrate(&datamodels.Product{})
	return
}

func (p *ProductManager) Insert(product *datamodels.Product) (productId int64, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	result := p.db.Create(product)
	return product.ID, result.Error
}

func (p *ProductManager) Delete(key int64) bool {
	if err := p.Conn(); err != nil {
		return false
	}

	result := p.db.Delete(&datamodels.Product{}, key)
	if result.Error != nil {
		log.Fatalln(result.Error)
		return false
	}
	return true
}

func (p *ProductManager) Update(product *datamodels.Product) (err error) {
	if err := p.Conn(); err != nil {
		return err
	}
	result := p.db.Save(product)
	return result.Error
}

func (p *ProductManager) SelectByKey(key int64) (product *datamodels.Product, err error) {
	product = &datamodels.Product{}

	if err = p.Conn(); err != nil {
		return product, err
	}
	result := p.db.First(&product, key)
	return product, result.Error
}

func (p *ProductManager) SelectAll() (products []*datamodels.Product, err error) {
	if err = p.Conn(); err != nil {
		return nil, err
	}

	result := p.db.Find(&products)
	return products, result.Error
}
