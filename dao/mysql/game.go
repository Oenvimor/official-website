package mysql

import (
	"cqupt_hub/models"
)

func CheckGameExist(gamePath string) (int64, error) {
	var count int64
	result := db.Model(&models.Game{}).Where("game_path = ?", gamePath).Count(&count)
	return count, result.Error
}

func InsertGame(g *models.Game) error {
	if err := db.Create(&g).Error; err != nil {
		return err
	}
	return nil
}

func EditGame(ID int, UpdateFields map[string]interface{}) error {
	if err := db.Model(models.Game{}).Where("id = ?", ID).Updates(UpdateFields).Error; err != nil {
		return err
	}
	return nil
}

func DeleteGame(ID int) error {
	err := db.Where("id = ?", ID).Delete(&models.Game{}).Error
	if err != nil {
		return err
	}
	return nil
}

func GetGame() (interface{}, error) {
	p := new([]models.Game)
	results := db.Table("games").Find(&p)
	if err := results.Error; err != nil {
		return nil, err
	}
	return p, nil
}

func GetImageUrl(ID int) (string, error) {
	p := new(models.Game)
	if err := db.Table("games").Select("game_image").Where("id = ?", ID).First(&p).Error; err != nil {
		return "", err
	}
	return p.GameImage, nil
}

func DisplayGame(fileID []int) error {
	if err := db.Table("games").Where("is_display", true).Update("is_display", false).Error; err != nil {
		return err
	}
	if err := db.Table("games").Where("id IN ?", fileID).Update("is_display", true).Error; err != nil {
		return err
	}
	return nil
}
