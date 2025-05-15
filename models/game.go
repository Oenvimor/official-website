package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	GameName  string `json:"gameName"`
	GamePath  string `json:"gamePath"`
	GameImage string `json:"gameImage"`
	IsDisplay bool   `json:"isDisplay"`
}
