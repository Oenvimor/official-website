package mysql

import (
	"cqupt_hub/models"
)

func CheckProjectExist(projectPath string) (int64, error) {
	var count int64
	result := db.Model(&models.Project{}).Where("project_path = ?", projectPath).Count(&count)
	return count, result.Error
}

func InsertProject(m *models.Project) error {
	if err := db.Create(&m).Error; err != nil {
		return err
	}
	return nil
}

func EditProject(ID int, UpdateFields map[string]interface{}) error {
	if err := db.Model(models.Project{}).Where("id = ?", ID).Updates(UpdateFields).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProject(ID int) error {
	err := db.Where("id = ?", ID).Delete(&models.Project{}).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProject() (interface{}, error) {
	p := new([]models.Project)
	results := db.Table("projects").Find(&p)
	if err := results.Error; err != nil {
		return nil, err
	}
	return p, nil
}

func GetImageUrlByID(ID int) (string, error) {
	p := new(models.Project)
	if err := db.Table("projects").Select("project_image").Where("id = ?", ID).First(&p).Error; err != nil {
		return "", err
	}
	return p.ProjectImage, nil
}

func DisplayProject(fileID []int) error {
	if err := db.Table("projects").Where("is_display", true).Update("is_display", false).Error; err != nil {
		return err
	}
	if err := db.Table("projects").Where("id IN ?", fileID).Update("is_display", true).Error; err != nil {
		return err
	}
	return nil
}
