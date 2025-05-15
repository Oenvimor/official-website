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
	ErrorProjectExist = errors.New("项目已存在")
)

func AddProject(m *models.Project) (err error) {
	// 验证项目是否存在
	var count int64
	count, err = mysql.CheckProjectExist(m.ProjectPath)
	if err != nil {
		fmt.Println("mysql.CheckProjectExist err - ", err)
		return
	}
	if count > 0 {
		return ErrorProjectExist
	}
	// 将数据插入数据库
	if err = mysql.InsertProject(m); err != nil {
		return err
	}
	return nil
}

func EditProject(ID int, UpdateFields map[string]interface{}) error {
	if len(UpdateFields) == 0 {
		return ErrorNoChange
	}
	if err := mysql.EditProject(ID, UpdateFields); err != nil {
		return err
	}
	return nil
}

func DeleteProject(ID int, accessKey string, secretKey string, bucket string) error {
	// 从数据库中获取头像信息
	url, err := mysql.GetImageUrlByID(ID)
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
	if err = mysql.DeleteProject(ID); err != nil {
		fmt.Println("delete from database error -", err)
		return err
	}
	return nil
}

func GetProject() (interface{}, error) {
	data, err := mysql.GetProject()
	if err != nil {
		fmt.Println("mysql.GetProject err -", err)
		return nil, err
	}
	return data, nil

}

func GetImageUrlByID(ID int) (string, error) {
	url, err := mysql.GetImageUrlByID(ID)
	if err != nil {
		fmt.Println("mysql.GetImageUrlByID err -", err)
		return "", err
	}
	return url, nil
}

func DisplayProject(fileID []int) error {
	err := mysql.DisplayProject(fileID)
	if err != nil {
		fmt.Println("set display project failed, err -", err)
		return err
	}
	return nil
}
