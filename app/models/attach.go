package models

import (
	"github.com/gookit/goutil/fmtutil"
	"github.com/goravel/framework/support/carbon"
)

const (
	TypeImage = "image"
	TypeFile  = "file"
	TypeVideo = "video"
	TypeAudio = "audio"
)

type Attach struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	CategoryID int    `gorm:"column:category_id" json:"categoryID"`
	Type       string `gorm:"column:type" json:"type"`
	Disk       string `gorm:"column:disk" json:"disk"`
	Name       string `gorm:"column:name" json:"name"`
	Mime       string `gorm:"column:mime" json:"mime"`
	Ext        string `gorm:"column:ext" json:"ext"`
	Size       int64  `gorm:"column:size" json:"size"`
	Path       string `gorm:"column:path" json:"path"`
	Url        string `gorm:"column:url" json:"url"`
	Timestamps
}

func (m *Attach) TableName() string {
	return "sys_attach"
}

type AttachResource struct {
	ID         int             `json:"id"`
	CategoryID int             `json:"categoryID"`
	Type       string          `json:"type"`
	Disk       string          `json:"disk"`
	Name       string          `json:"name"`
	Mime       string          `json:"mime"`
	Ext        string          `json:"ext"`
	Size       string          `json:"size"`
	Path       string          `json:"path"`
	Url        string          `json:"url"`
	CreatedAt  carbon.DateTime `json:"createdAt"`
}

func (m *Attach) Resource() AttachResource {
	disk := ""
	switch m.Disk {
	case "local":
		disk = "本地"
	case "oss":
		disk = "阿里云"
	case "cos":
		disk = "腾讯云"
	}
	return AttachResource{
		ID:         m.ID,
		CategoryID: m.CategoryID,
		Disk:       disk,
		Name:       m.Name,
		Mime:       m.Mime,
		Url:        m.Url,
		Size:       fmtutil.DataSize(uint64(m.Size)),
		CreatedAt:  m.CreatedAt,
	}
}

type AttachCategory struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	ParentID int    `json:"parentID"`
	Name     string `json:"name"`
	Sort     int    `json:"sort"`
	Remark   string `json:"remark"`
	IsEnable int    `json:"isEnable"`
	Timestamps

	// 关联表
	Attaches []*Attach `json:"attaches,omitempty" gorm:"foreignKey:CategoryID"`
}

func (m *AttachCategory) TableName() string {
	return "sys_attach_category"
}
