package logic

import (
	"cqupt_hub/dao/mysql"
	"cqupt_hub/models"
	"errors"
	"fmt"
)

var ErrorPositionExist = errors.New("部门已存在")

func AddPosition(p *models.Position) error {
	// 检查部门是否存在
	count, err := mysql.CheckPositionExist(p.PositionName)
	if err != nil {
		fmt.Println("check position err -", err)
		return err
	}
	if count > 0 {
		return ErrorPositionExist
	}
	// 将数据插入数据库
	if err = mysql.InsertPosition(p); err != nil {
		fmt.Println("insert position err -", err)
		return err
	}
	return nil
}

func GetPosition() (interface{}, error) {
	data, err := mysql.GetPosition()
	if err != nil {
		fmt.Println("get position err -", err)
		return nil, err
	}
	return data, nil
}

func EditPosition(ID int, UpdateFields map[string]interface{}) error {
	if len(UpdateFields) == 0 {
		return ErrorNoChange
	}
	if err := mysql.EditPosition(ID, UpdateFields); err != nil {
		return err
	}
	return nil
}

func DeletePosition(ID int) error {
	if err := mysql.DeletePosition(ID); err != nil {
		return err
	}
	return nil
}
