package models

import "gorm.io/gorm"

// Department 部长信息
type Department struct {
	gorm.Model
	SectionName  string `json:"departName"`
	Avatar       string `json:"personAvatar"`
	Username     string `json:"personName"`
	Email        string `json:"email"`
	Introduction string `json:"personIntro"`
}
