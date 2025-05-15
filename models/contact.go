package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	StudioAddress string `json:"studioAddr"`
	Email         string `json:"email"`
	PhoneNum      string `json:"phoneNum"`
	QQ            string `json:"qq"`
}
