package models

import "gorm.io/gorm"

type ParamRegister struct {
	gorm.Model
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"rePassword"`
}
