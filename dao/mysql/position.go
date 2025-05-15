package mysql

import (
	"cqupt_hub/models"
	"fmt"
)

func CheckPositionExist(positionName string) (count int64, err error) {
	result := db.Table("positions").Where("position_name = ? AND deleted_at IS NULL", positionName).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func InsertPosition(p *models.Position) error {
	if err := db.Create(&p).Error; err != nil {
		fmt.Println("insert position error -", err)
		return err
	}
	return nil
}

func GetPosition() (interface{}, error) {
	p := new([]models.Position)
	results := db.Table("positions").Find(&p)
	if err := results.Error; err != nil {
		return nil, err
	}
	return p, nil
}

func EditPosition(ID int, UpdateFields map[string]interface{}) error {
	if err := db.Model(models.Position{}).Where("id = ?", ID).Updates(UpdateFields).Error; err != nil {
		return err
	}
	return nil
}

func DeletePosition(ID int) error {
	if err := db.Where("id = ?", ID).Delete(&models.Position{}).Error; err != nil {
		return err
	}
	return nil
}
