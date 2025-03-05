package models

import (
	"github.com/goravel/framework/support/carbon"
)

const (
	IsEnable = 1
)

type Timestamps struct {
	CreatedAt carbon.DateTime `gorm:"autoCreateTime;column:created_at" json:"createdAt"`
	UpdatedAt carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updatedAt"`
}
