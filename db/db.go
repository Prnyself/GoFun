package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Connect() (*xorm.Engine, error) {
	return xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/func?charset=utf8")
}