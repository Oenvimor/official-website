package logic

import (
	"cqupt_hub/dao/mysql"
	"cqupt_hub/models"
)

func AddContact(contact *models.Contact) error {
	// 将数据插入数据库
	if err := mysql.AddContact(contact); err != nil {
		return err
	}
	return nil
}

func GetContact() (interface{}, error) {
	data, err := mysql.GetContact()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func EditContact(updateFields map[string]interface{}) error {
	if err := mysql.EditContact(updateFields); err != nil {
		return err
	}
	return nil
}
