package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/xiaoheigou/GoFun/config"
)

func Connect() (*xorm.Engine, error) {
	fmt.Sprintf("%s:%d@tcp(localhost:3306)/%s?charset=utf8", config.USER_NAME, config.USER_PWD, config.DATABASE)
	return xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/func?charset=utf8")
}
func Find(DB *xorm.Engine, findBy interface{}, objects interface{}) error {
	return DB.Find(objects, findBy)
}

func FindBy(DB *xorm.Engine, model interface{}) (err error) {
	_, err = DB.Get(model)
	return
}

func Exists(DB *xorm.Engine, model interface{}) (bool, error) {
	return DB.Get(model)
}

func Update(DB *xorm.Engine, id int64, model interface{}) (err error) {
	_, err = DB.Id(id).Update(model)
	return
}

func Store(DB *xorm.Engine, model interface{}) (err error) {
	_, err = DB.Insert(model)
	return
}

func Destroy(DB *xorm.Engine, id int64, model interface{}) (err error) {
	_, err = DB.Id(id).Delete(model)
	return
}
