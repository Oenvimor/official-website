package models

import "gorm.io/gorm"

// ImageRecord 图片管理
type ImageRecord struct {
	gorm.Model
	FileName   string `json:"file_name"`   // 文件名
	URL        string `json:"url"`         // 图片的访问 URL
	IsHomepage bool   `json:"is_homepage"` // 是否被用作首页图片
}
