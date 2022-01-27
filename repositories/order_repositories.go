package repositories

import (
	"eshop/common"
	"eshop/datamodels"
	"log"

	"gorm.io/gorm"
)

type IOrder interface {
	Conn() error
	Insert(*datamodels.Order) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Order) error
	SelectByKey(int64) (*datamodels.Order, error)
	SelectAll() ([]*datamodels.Order, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}

type OrderManager struct {
	table string
	db    *gorm.DB
}

type OrderInfoResult struct {
	orderID     int64
	productName string
	orderStatus int64
}

func NewOrderManager(table string, db *gorm.DB) IOrder {
	return &OrderManager{table: table, db: db}
}

func (o *OrderManager) Conn() (err error) {
	if o.db == nil {
		db, err := common.GetNewGormDB()
		if err != nil {
			return err
		}
		o.db = db
	}
	if o.table == "" {
		o.table = "order"
	}
	o.db.AutoMigrate(&datamodels.Order{})
	return
}

func (o *OrderManager) Insert(order *datamodels.Order) (orderId int64, err error) {
	if err = o.Conn(); err != nil {
		return
	}

	result := o.db.Create(order)
	return order.ID, result.Error
}

func (o *OrderManager) Delete(key int64) bool {
	if err := o.Conn(); err != nil {
		return false
	}

	result := o.db.Delete(&datamodels.Order{}, key)
	if result.Error != nil {
		log.Fatal("delete order failed", result.Error)
		return false
	}
	return true
}

func (o *OrderManager) Update(order *datamodels.Order) error {
	if err := o.Conn(); err != nil {
		return err
	}

	result := o.db.Save(order)
	return result.Error
}

func (o *OrderManager) SelectByKey(key int64) (order *datamodels.Order, err error) {
	order = &datamodels.Order{}
	if err = o.Conn(); err != nil {
		return order, err
	}
	result := o.db.First(&order, key)
	return order, result.Error
}

func (o *OrderManager) SelectAll() (orders []*datamodels.Order, err error) {
	if err = o.Conn(); err != nil {
		return nil, err
	}

	result := o.db.Find(&orders)
	return orders, result.Error
}

func (o *OrderManager) SelectAllWithInfo() (orders map[int]map[string]string, err error) {
	if err = o.Conn(); err != nil {
		return nil, err
	}

	rows, err := o.db.Table("orders").Select(
		"orders.id, products.product_name, orders.order_status").Joins(
		"left join products on orders.id = products.id").Rows()

	orders = common.GetResultRows(rows)
	return orders, err
}
