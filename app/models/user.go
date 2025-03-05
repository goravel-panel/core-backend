package models

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Name   string `gorm:"column:name"`
	Avatar string `gorm:"column:avatar"`
	orm.SoftDeletes
}

func (m *User) TableName() string {
	return "user"
}
