package models

import (
	"github.com/goravel/framework/support/carbon"
)

type Article struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	CategoryID int    `gorm:"column:category_id" json:"categoryID"`
	From       string `gorm:"column:from" json:"from"`
	Title      string `gorm:"column:title" json:"title"`
	Cover      string `gorm:"column:cover" json:"cover"`
	Content    string `gorm:"column:content" json:"content"`
	IsEnable   int    `gorm:"column:is_enable" json:"isEnable"`
	Timestamps

	// 关联表
	Category *ArticleCategory `json:"category"`
}

func (m *Article) TableName() string {
	return "article"
}

func (m *Article) Resource() any {
	return struct {
		ID         int             `gorm:"primaryKey" json:"id"`
		CategoryID int             `json:"categoryID"`
		From       string          `json:"from"`
		Title      string          `json:"title"`
		Cover      string          `json:"cover"`
		Content    string          `json:"content"`
		IsEnable   int             `json:"isEnable"`
		CreatedAt  carbon.DateTime `json:"createdAt"`
	}{
		ID:         m.ID,
		CategoryID: m.CategoryID,
		From:       m.From,
		Title:      m.Title,
		Cover:      m.Cover,
		Content:    m.Content,
		IsEnable:   m.IsEnable,
		CreatedAt:  m.CreatedAt,
	}
}

type ArticleCategory struct {
	ID        int             `gorm:"primaryKey" json:"id"`
	Name      string          `gorm:"column:name" json:"name"`
	Sort      int             `gorm:"column:sort" json:"sort"`
	Remark    string          `gorm:"column:remark" json:"remark"`
	IsEnable  int             `gorm:"column:is_enable" json:"isEnable"`
	CreatedAt carbon.DateTime `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updatedAt"`

	ArticleCount int64 `gorm:"-" json:"articleCount"`
}

func (m *ArticleCategory) TableName() string {
	return "article_category"
}
