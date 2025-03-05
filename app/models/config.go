package models

const (
	IsDefault = 1

	TypeBase    = "base"
	TypeCaptcha = "captcha"
	TypeMail    = "mail"
	TypeStorage = "storage"
	TypePage    = "page"
)

type Config struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Type      string `gorm:"column:type" json:"type"`
	Name      string `gorm:"column:name" json:"name"`
	Title     string `gorm:"column:title" json:"title"`
	Value     string `gorm:"type:json;column:value" json:"-"`
	Remark    string `gorm:"column:remark" json:"remark"`
	IsDefault int    `gorm:"column:is_default" json:"isDefault"`
	IsEnable  int    `gorm:"column:is_enable" json:"isEnable"`
	Sort      int    `gorm:"column:sort" json:"sort"`
	Timestamps

	// 关联表
	Attaches []*Attach `json:"attaches" gorm:"foreignKey:Disk;references:Name"`
}

func (*Config) TableName() string {
	return "sys_config"
}
