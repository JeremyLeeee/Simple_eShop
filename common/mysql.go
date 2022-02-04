package common

import (
	"database/sql"
	"log"
	"reflect"
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var CONNPATH = "root:iam59!z$@tcp(106.52.119.98:3306)/eshop"
var once sync.Once
var Db *sql.DB

func NewMysqlConn() (db *sql.DB, err error) {
	// singleton
	once.Do(func() {
		Db, err = sql.Open("mysql", CONNPATH)
		if err != nil {
			log.Println(err)
		}
	})

	return Db, err
}

// get one row
func GetResultRow(rows *sql.Rows) map[string]string {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}
	record := make(map[string]string)
	for rows.Next() {
		rows.Scan(scanArgs...)
		for i, v := range values {
			if v != nil {
				if reflect.TypeOf(v).Name() == "int64" {
					record[columns[i]] = strconv.FormatInt(v.(int64), 10)
				} else {
					record[columns[i]] = string(v.([]byte))
				}

			}
		}
	}
	return record
}

// get all rows
func GetResultRows(rows *sql.Rows) map[int]map[string]string {
	columns, _ := rows.Columns()
	vals := make([][]byte, len(columns))
	scans := make([]interface{}, len(columns))
	for k := range vals {
		scans[k] = &vals[k]
	}
	i := 0
	result := make(map[int]map[string]string)
	for rows.Next() {
		rows.Scan(scans...)
		row := make(map[string]string)
		for k, v := range vals {
			key := columns[k]
			row[key] = string(v)
		}
		result[i] = row
		i++
	}
	return result
}
