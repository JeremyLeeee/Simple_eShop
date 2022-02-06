package repositories

import (
	"database/sql"
	"eshop/common"
	"eshop/datamodels"
	"log"

	"strconv"
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
	table     string
	mysqlConn *sql.DB
}

func NewOrderManager(table string, sql *sql.DB) IOrder {
	return &OrderManager{table: table, mysqlConn: sql}
}

func (o *OrderManager) Conn() (err error) {
	if o.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		o.mysqlConn = mysql
	}
	if o.table == "" {
		o.table = "eshop.order"
	}
	return
}

func (o *OrderManager) Insert(order *datamodels.Order) (orderId int64, err error) {
	if err = o.Conn(); err != nil {
		return
	}

	sql := "INSERT " + o.table + " SET ID=?, userId=?, productID=?, orderStatus=?"
	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	result, err := stmt.Exec(order.ID, order.UserId, order.ProductId, order.OrderStatus)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return result.LastInsertId()
}

func (o *OrderManager) Delete(key int64) bool {
	if err := o.Conn(); err != nil {
		return false
	}

	sql := "DELETE FROM " + o.table + " WHERE ID=?"
	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		log.Println(err)
		return false
	}

	_, err = stmt.Exec(key)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (o *OrderManager) Update(order *datamodels.Order) error {
	if err := o.Conn(); err != nil {
		return err
	}

	sql := "UPDATE " + o.table + " SET ID=?, userId=?, productID=?, orderStatus=?, WHERE ID=?"
	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(order.ID, order.ProductId, order.UserId,
		order.OrderStatus, strconv.FormatInt(order.ID, 10))

	if err != nil {
		return err
	}

	return nil
}

func (o *OrderManager) SelectByKey(key int64) (order *datamodels.Order, err error) {
	order = &datamodels.Order{}
	if err = o.Conn(); err != nil {
		return order, err
	}
	sql := "SELECT * from " + o.table + " WHERE ID = " + strconv.FormatInt(key, 10)
	row, err := o.mysqlConn.Query(sql)
	if err != nil {
		return order, err
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return order, err
	}
	common.DataToStructByTagSql(result, order)
	return
}

func (o *OrderManager) SelectAll() (orders []*datamodels.Order, err error) {
	if err = o.Conn(); err != nil {
		return nil, err
	}

	sql := "SELECT * FROM " + o.table
	rows, err := o.mysqlConn.Query(sql)
	// defer rows.Close()

	if err != nil {
		return nil, err
	}
	results := common.GetResultRows(rows)
	if len(results) == 0 {
		return nil, err
	}

	for _, v := range results {
		order := &datamodels.Order{}
		common.DataToStructByTagSql(v, order)
		orders = append(orders, order)
	}

	return
}

func (o *OrderManager) SelectAllWithInfo() (orders map[int]map[string]string, err error) {
	if err = o.Conn(); err != nil {
		return nil, err
	}

	sql := "SELECT o.ID, o.userId, p.productName,o.orderStatus " +
		"FROM eshop.order AS o LEFT JOIN product AS p ON o.productID=p.ID"

	rows, err := o.mysqlConn.Query(sql)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	orders = common.GetResultRows(rows)
	return orders, err
}
