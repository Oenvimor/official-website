package models

import "gorm.io/gorm"

// Position 岗位信息
type Position struct {
	gorm.Model
	BelongTo     string `json:"belong_to,omitempty"`
	PositionName string `json:"positionName,omitempty"`
	Requirement  string `json:"requirement,omitempty"`
	Delivery     string `json:"deliveryEmail,omitempty"`
}
