package models

type Role struct {
	ID       int    `gorm:"primaryKey;id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Remark   string `gorm:"column:remark" json:"remark"`
	Sort     int    `gorm:"column:sort" json:"sort"`
	IsEnable int    `gorm:"column:is_enable" json:"isEnable"`
	Timestamps

	// 关联表
	Menus   []*Menu  `gorm:"many2many:sys_role_menu" json:"menus"`
	Members []*Admin `gorm:"many2many:sys_admin_role" json:"members"`
}

func (*Role) TableName() string {
	return "sys_role"
}

type RoleMenu struct {
	RoleID int `json:"roleID"`
	MenuID int `json:"menuID"`
}

func (*RoleMenu) TableName() string {
	return "sys_role_menu"
}
