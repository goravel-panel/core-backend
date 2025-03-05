package models

const (
	StatusSuccess = "success"
	StatusFail    = "fail"
)

type LoginLogger struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	AdminID  int    `gorm:"column:admin_id" json:"adminID"`
	Account  string `gorm:"column:account" json:"account"`
	IP       string `gorm:"column:ip" json:"ip"`
	Location string `gorm:"column:location" json:"location"`
	Browser  string `gorm:"column:browser" json:"browser"`
	OS       string `gorm:"column:os" json:"os"`
	Status   string `gorm:"column:status" json:"status"`
	Timestamps
}

func (m *LoginLogger) TableName() string {
	return "sys_login_log"
}
