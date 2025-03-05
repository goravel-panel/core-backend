package models

import (
	"github.com/goravel/framework/support/carbon"
)

const (
	IsRoot = 1
)

type Admin struct {
	ID            int             `gorm:"primaryKey" json:"id"`
	Account       string          `gorm:"column:account" json:"account"`
	Password      string          `gorm:"column:password" json:"-"`
	Name          string          `gorm:"column:name" json:"name"`
	Avatar        string          `gorm:"column:avatar" json:"avatar"`
	IsEnable      int             `gorm:"column:is_enable" json:"isEnable"`
	IsRoot        int             `gorm:"column:is_root" json:"isRoot"`
	LastLoginIP   string          `gorm:"column:last_login_ip" json:"lastLoginIP"`
	LastLoginTime carbon.DateTime `gorm:"column:last_login_time" json:"lastLoginTime"`
	Timestamps

	// 关联表
	Roles []*Role `gorm:"many2many:sys_admin_role" json:"roles"`
}

func (m *Admin) TableName() string {
	return "sys_admin"
}

type AdminRole struct {
	AdminID int `json:"adminID"`
	RoleID  int `json:"roleID"`
}

func (*AdminRole) TableName() string {
	return "sys_admin_role"
}
