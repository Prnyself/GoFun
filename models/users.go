package users

import "github.com/xiaoheigou/GoFun/pkg/db"

type Users []User
type User db.Users

func (u *User) TableName() string {
	return "users"
}
