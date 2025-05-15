package logic

import (
	"cqupt_hub/dao/mysql"
	"cqupt_hub/models"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
)

var (
	ErrorUserExist = errors.New("用户已存在")
	ErrorNoChange  = errors.New("未作修改")
)

func AddDepartment(m *models.Department) (err error) {
	// 验证用户是否存在
	var count int64
	count, err = mysql.CheckDepartmentExist(m.SectionName)
	if err != nil {
		fmt.Println("mysql.CheckUserExist err - ", err)
		return
	}
	if count > 0 {
		return ErrorUserExist
	}
	// 将数据插入数据库
	if err = mysql.InsertDepartment(m); err != nil {
		return err
	}
	return nil
}

func EditDepartment(UserID int, UpdateFields map[string]interface{}) error {
	if len(UpdateFields) == 0 {
		return ErrorNoChange
	}
	if err := mysql.EditDepartment(UserID, UpdateFields); err != nil {
		return err
	}
	return nil
}

func DeleteDepartment(UserID int, accessKey string, secretKey string, bucket string) error {
	// 从数据库中获取头像信息
	url, err := mysql.GetAvatarUrlByID(UserID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Url := strings.Split(url, "/")
			// 从千牛删除图片
			err = QiniuDelete(Url[1], accessKey, secretKey, bucket)
			if err != nil {
				fmt.Println("QiniuDelete err -", err)
				return err
			}
		} else {
			fmt.Println("mysql.GetAvatarUrlByID err -", err)
			return err
		}
	}
	// 从数据库中删除部门信息
	if err = mysql.DeleteDepartment(UserID); err != nil {
		fmt.Println("delete from database error -", err)
		return err
	}
	return nil
}

func GetDepartment() (interface{}, error) {
	data, err := mysql.GetDepartment()
	if err != nil {
		fmt.Println("mysql.GetDepartment err -", err)
		return nil, err
	}
	return data, nil

}

func GetAvatarUrlByID(ID int) (string, error) {
	url, err := mysql.GetAvatarUrlByID(ID)
	if err != nil {
		return "", err
	}
	return url, nil
}

func GetSectionNameByID(userID int) (string, error) {
	Name, err := mysql.GetSectionNameByID(userID)
	if err != nil {
		return "", err
	}
	return Name, nil
}

func DeletePositionBySectionName(Name string) error {
	if err := mysql.DeletePositionBySectionName(Name); err != nil {
		return err
	}
	return nil
}
