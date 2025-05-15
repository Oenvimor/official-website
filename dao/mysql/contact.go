package mysql

import "cqupt_hub/models"

func AddContact(contact *models.Contact) error {
	if err := db.Create(&contact).Error; err != nil {
		return err
	}
	return nil
}

func GetContact() (interface{}, error) {
	var contact *models.Contact
	if err := db.Model(&models.Contact{}).Find(&contact).Error; err != nil {
		return nil, err
	}
	return contact, nil
}

func EditContact(updateFields map[string]interface{}) error {
	if err := db.Model(&models.Contact{}).Where("id = ?", 1).Updates(updateFields).Error; err != nil {
		return err
	}
	return nil
}
