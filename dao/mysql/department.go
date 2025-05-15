package mysql

import (
	"cqupt_hub/models"
)

func CheckDepartmentExist(section string) (int64, error) {
	var count int64
	result := db.Model(&models.Department{}).Where("section_name = ?", section).Count(&count)
	return count, result.Error
}

func InsertDepartment(m *models.Department) error {
	if err := db.Create(&m).Error; err != nil {
		return err
	}
	return nil
}

func EditDepartment(UserID int, UpdateFields map[string]interface{}) error {
	if err := db.Model(models.Department{}).Where("id = ?", UserID).Updates(UpdateFields).Error; err != nil {
		return err
	}
	return nil
}

func GetAvatarUrlByID(UserID int) (string, error) {
	i := new(models.Department)
	if err := db.Table("departments").Select("avatar").Where("id = ?", UserID).First(&i).Error; err != nil {
		return "", err
	}
	return i.Avatar, nil
}

func DeleteDepartment(UserID int) error {
	err := db.Where("id = ?", UserID).Delete(&models.Department{}).Error
	if err != nil {
		return err
	}
	return nil
}

func GetDepartment() (interface{}, error) {
	d := new([]models.Department)
	results := db.Table("departments").Find(&d)
	if err := results.Error; err != nil {
		return nil, err
	}
	return d, nil
}

func GetSectionNameByID(UserID int) (string, error) {
	var request struct {
		Name string
	}
	result := db.Model(&models.Department{}).Select("section_name").Where("id = ?", UserID).Scan(&request)
	if err := result.Error; err != nil {
		return "", err
	}
	return request.Name, nil
}

func DeletePositionBySectionName(Name string) error {
	if err := db.Where("belong_to = ?", Name).Delete(&models.Position{}).Error; err != nil {
		return err
	}
	return nil
}
