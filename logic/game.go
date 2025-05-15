package logic

import (
	"cqupt_hub/dao/mysql"
	"cqupt_hub/models"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

var (
	ErrorGameExist = errors.New("游戏已存在")
)

func AddGame(g *models.Game) (err error) {
	// 验证项目是否存在
	var count int64
	count, err = mysql.CheckGameExist(g.GamePath)
	if err != nil {
		fmt.Println("mysql.CheckProjectExist err - ", err)
		return
	}
	if count > 0 {
		return ErrorGameExist
	}
	// 将数据插入数据库
	if err = mysql.InsertGame(g); err != nil {
		return err
	}
	return nil
}

func EditGame(ID int, UpdateFields map[string]interface{}) error {
	if len(UpdateFields) == 0 {
		return ErrorNoChange
	}
	if err := mysql.EditGame(ID, UpdateFields); err != nil {
		return err
	}
	return nil
}

func DeleteGame(ID int, accessKey string, secretKey string, bucket string) error {
	// 从数据库中获取头像信息
	url, err := mysql.GetImageUrl(ID)
	if err != nil {
		fmt.Println("mysql.GetImageUrl err -", err)
		return err
	}
	Url := strings.Split(url, "/")
	// 从千牛删除图片
	err = QiniuDelete(Url[1], accessKey, secretKey, bucket)
	if err != nil {
		fmt.Println("QiniuDelete err -", err)
		return err
	}
	// 从数据库中删除部门信息
	if err = mysql.DeleteGame(ID); err != nil {
		fmt.Println("delete from database error -", err)
		return err
	}
	return nil
}

func GetGame() (interface{}, error) {
	data, err := mysql.GetGame()
	if err != nil {
		fmt.Println("mysql.GetGame err -", err)
		return nil, err
	}
	return data, nil

}

func DisplayGame(fileID []int) error {
	err := mysql.DisplayGame(fileID)
	if err != nil {
		fmt.Println("set display game failed, err -", err)
		return err
	}
	return nil
}
func GetImageUrl(ID int) (string, error) {
	url, err := mysql.GetImageUrl(ID)
	if err != nil {
		return "", err
	}
	return url, nil
}
