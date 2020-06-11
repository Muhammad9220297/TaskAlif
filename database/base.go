package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysqlDB *gorm.DB
var err error

//Connect...
func Connect() error {
	mysqlDB, err = gorm.Open("mysql", "root:@/alifsender")
	return err
}

//DB_J...
func DB_J() *gorm.DB {
	return mysqlDB
}
