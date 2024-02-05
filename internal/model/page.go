// Code generated by sql2gorm. DO NOT EDIT.
package model

import (
	"time"
)

type Page struct {
	ID          uint      `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	AbsoluteUrl string    `gorm:"column:absolute_url" json:"absolute_url"` // 完整链接
	Status      int       `gorm:"column:status;default:0" json:"status"`   // 0-待爬取，1-已爬取
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (m *Page) TableName() string {
	return "page"
}
