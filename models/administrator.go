package models

import (
	"gorm.io/gorm"
)

// Administrator 管理员信息
type Administrator struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
