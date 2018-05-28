package v1

import (
	"github.com/go-xorm/xorm"
	"github.com/xiaoheigou/GoFun/pkg/db"
)

var dd *xorm.Engine

func init() {
	dd, _ = db.Connect()
}
