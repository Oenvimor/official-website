package logic

import (
	"cqupt_hub/dao/mysql"
	"cqupt_hub/models"
	"cqupt_hub/pkg/jwt"
	"fmt"
	"github.com/pkg/errors"
)

var (
	ErrorUserNotExist       = errors.New("用户不存在")
	ErrorWrongPassword      = errors.New("密码错误")
	ErrorSamePassword       = errors.New("新密码不能和旧密码相同")
	ErrorTwoPasswordNotSame = errors.New("两次密码输入不相同")
)

func Register(r *models.ParamRegister) (err error) {
	// 验证用户是否存在
	var count int64
	count, err = mysql.CheckUserExist(r.Username)
	if err != nil {
		fmt.Println("mysql.CheckUserExist err - ", err)
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	// 将数据插入数据库
	admin := &models.Administrator{
		Username: r.Username,
		Password: r.Password,
	}
	if err = mysql.InsertUser(admin); err != nil {
		return err
	}
	return nil
}

func Login(r *models.ParamRegister) (token string, err error) {
	// 验证用户是否存在
	var count int64
	count, err = mysql.CheckUserExist(r.Username)
	if err != nil {
		fmt.Println("mysql.CheckUserExist err - ", err)
		return
	}
	if count == 0 {
		return "", ErrorUserNotExist
	}
	// 验证密码是否正确
	var password string
	password, err = mysql.GetPasswordByUsername(r.Username)
	if err != nil {
		fmt.Println("mysql.PasswordByUsername err - ", err)
		return
	}
	if password != r.Password {
		return "", ErrorWrongPassword
	}
	// 生成token
	token, err = jwt.GenToken(r.Username, int64(r.ID))
	if err != nil {
		fmt.Println("jwt.GenToken err - ", err)
		return
	}
	return token, nil
}

func EditPassword(username string, password string, newPassword string, reNewPassword string) error {
	// 校验参数
	// 验证用户是否存在
	count, err := mysql.CheckUserExist(username)
	if err != nil {
		fmt.Println("mysql.CheckUserExist err - ", err)
		return err
	}
	if count == 0 {
		return ErrorUserNotExist
	}
	// 验证用户原密码是否正确
	OPassword, err := mysql.GetPasswordByUsername(username)
	if err != nil {
		fmt.Println("mysql.GetPasswordByUsername err - ", err)
		return err
	}
	if password != OPassword {
		return ErrorWrongPassword
	}
	// 验证新密码和旧密码是否相同
	if newPassword == OPassword {
		return ErrorSamePassword
	}
	// 验证新密码的两次输入是否相同
	if newPassword != reNewPassword {
		return ErrorTwoPasswordNotSame
	}
	// 更新数据库
	if err = mysql.EditPassword(username, newPassword); err != nil {
		return err
	}
	return nil
}
