package common

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetNewGormDB() (*gorm.DB, error) {
	gormdb, err := gorm.Open(mysql.Open(CONNPATH))
	if err != nil {
		log.Fatalln(err)
	}
	return gormdb, err
}
