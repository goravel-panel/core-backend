package models

const (
	IsShow = 1

	TypeCatalogue = "catalogue"
	TypeMenu      = "menu"
	TypeButton    = "button"
	TypeLink      = "link"
	TypeIframe    = "iframe"
)

type Menu struct {
	ID           int    `gorm:"primaryKey;" json:"id"`
	ParentID     int    `gorm:"column:parent_id" json:"pid"`
	Type         string `gorm:"column:type" json:"type"`
	Name         string `gorm:"column:name" json:"name"`
	Icon         string `gorm:"column:icon" json:"icon"`
	Sort         int    `gorm:"column:sort" json:"sort"`
	RoutePath    string `gorm:"column:route_path" json:"routePath"`
	SelectedPath string `gorm:"column:selected_path" json:"selectedPath"`
	Url          string `gorm:"column:url" json:"url"`
	Param        string `gorm:"column:param" json:"param"`
	Perm         string `gorm:"column:perm" json:"perm"`
	Component    string `gorm:"column:component" json:"component"`
	IsCache      int    `gorm:"is_cache" json:"-"`
	IsShow       int    `gorm:"column:is_show" json:"isShow"`
	IsEnable     int    `gorm:"column:is_enable" json:"isEnable"`
	Timestamps
}

type TreeMenu struct {
	Menu
	Children []TreeMenu `json:"children,omitempty" gorm:"-"`
}

func (*Menu) TableName() string {
	return "sys_menu"
}
