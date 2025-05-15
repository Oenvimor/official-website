package mysql

import "cqupt_hub/models"

func CheckUserExist(username string) (int64, error) {
	var count int64
	result := db.Model(&models.Administrator{}).Where("username = ?", username).Count(&count)
	return count, result.Error
}

func InsertUser(a *models.Administrator) error {
	if err := db.Create(&a).Error; err != nil {
		return err
	}
	return nil
}

func GetPasswordByUsername(username string) (string, error) {
	var password string
	result := db.Model(&models.Administrator{}).Where("username = ?", username).Pluck("password", &password)
	return password, result.Error
}

func EditPassword(username string, newPassword string) error {
	if err := db.Model(&models.Administrator{}).Where("username = ?", username).Update("password", newPassword).Error; err != nil {
		return err
	}
	return nil
}
