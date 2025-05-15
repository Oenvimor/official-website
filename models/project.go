package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ProjectName  string `json:"projectName"`
	ProjectPath  string `json:"projectPath"`
	ProjectImage string `json:"projectImage"`
	IsDisplay    bool   `json:"isDisplay"`
}
