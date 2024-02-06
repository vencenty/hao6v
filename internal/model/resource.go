package model

import (
	"time"
)

type Resource struct {
	ID            uint      `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	PageID        uint      `gorm:"column:page_id;default:0" json:"page_id"`     // 关联pages的id
	ResourceTitle string    `gorm:"column:resource_title" json:"resource_title"` // 下载资源标题
	DownloadUrl   string    `gorm:"column:download_url" json:"download_url"`     // 下载链接的url
	Type          string    `gorm:"column:type" json:"type"`                     // 资源类型
	CreatedAt     time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (m *Resource) TableName() string {
	return "resource"
}
