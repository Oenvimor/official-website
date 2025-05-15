package mysql

import (
	"cqupt_hub/models"
)

func Upload(record *models.ImageRecord) error {
	if err := db.Create(&record).Error; err != nil {
		return err
	}
	return nil
}

func Delete(filename string) error {
	if err := db.Where("file_name = ?", filename).Delete(&models.ImageRecord{}).Error; err != nil {
		return err
	}
	return nil
}

func GetImage() (records []models.ImageRecord, err error) {
	result := db.Table("image_records").Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}

func SetHomePage(fileName []string) error {
	if err := db.Table("image_records").Where("is_homepage = ?", true).Update("is_homepage", false).Error; err != nil {
		return err
	}
	if err := db.Table("image_records").Where("file_name IN ?", fileName).Update("is_homepage", true).Error; err != nil {
		return err
	}
	return nil
}

func CheckImageByName(filename string) (count int64, err error) {
	if err = db.Table("image_records").Where("file_name = ? AND deleted_at IS NULL", filename).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
