package repositories

import (
	"database/sql"
	"eshop/common"
	"eshop/datamodels"
	"log"
	"strconv"
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
	table     string
	mysqlConn *sql.DB
}

func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductManager{table: table, mysqlConn: db}
}

func (p *ProductManager) Conn() (err error) {
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = "product"
	}
	return
}

func (p *ProductManager) Insert(product *datamodels.Product) (productId int64, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	// define sql script and prepare
	sql := "INSERT product SET productName=?, productNum=?, productImage=?, productUrl=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	// execute sql
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return result.LastInsertId()
}

func (p *ProductManager) Delete(productId int64) bool {
	if err := p.Conn(); err != nil {
		return false
	}
	// define sql script and prepare
	sql := "DELETE FROM product WHERE ID=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		log.Println(err)
		return false
	}

	// execute sql
	_, err = stmt.Exec(productId)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (p *ProductManager) Update(product *datamodels.Product) (err error) {
	if err := p.Conn(); err != nil {
		return err
	}

	// define sql script and prepare
	sql := "UPDATE product SET productName=?, productNum=?, productImage=?, productUrl=? WHERE ID=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}

	// execute sql
	_, err = stmt.Exec(product.ProductName, product.ProductNum,
		product.ProductImage, product.ProductUrl, strconv.FormatInt(product.ID, 10))

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductManager) SelectByKey(key int64) (product *datamodels.Product, err error) {
	product = &datamodels.Product{}
	if err = p.Conn(); err != nil {
		return product, err
	}
	sql := "SELECT * from " + p.table + " WHERE ID = " + strconv.FormatInt(key, 10)
	row, err := p.mysqlConn.Query(sql)
	if err != nil {
		return product, err
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return product, err
	}
	common.DataToStructByTagSql(result, product)
	return
}

func (p *ProductManager) SelectAll() (products []*datamodels.Product, err error) {
	if err = p.Conn(); err != nil {
		return nil, err
	}

	sql := "SELECT * FROM " + p.table
	rows, err := p.mysqlConn.Query(sql)
	// defer rows.Close()

	if err != nil {
		return nil, err
	}
	results := common.GetResultRows(rows)
	if len(results) == 0 {
		return nil, nil
	}

	for _, v := range results {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v, product)
		// log.Print("range find: " + product.ProductName)
		products = append(products, product)

	}

	return
}
