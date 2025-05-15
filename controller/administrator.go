package controller

import (
	"cqupt_hub/logic"
	"cqupt_hub/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

var CtxUsernameKey = "username"

func RegisterHandler(c *gin.Context) {
	// 获取参数
	r := new(models.ParamRegister)
	if err := c.ShouldBindJSON(r); err != nil {
		fmt.Println("invalid request params -", err)
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 处理参数
	// 手动处理
	if r.Username == "" || r.Password == "" || r.RePassword == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "必填字段不能为空",
		})
		return
	}
	if r.Password != r.RePassword {
		c.JSON(http.StatusOK, gin.H{
			"msg": "两次输入密码不一致",
		})
		return
	}
	// 逻辑处理
	if err := logic.Register(r); err != nil {
		fmt.Println("logic.Register failed -", err)
		if errors.Is(err, logic.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// 获取参数
	l := new(models.ParamRegister)
	if err := c.ShouldBindJSON(l); err != nil {
		fmt.Println("invalid request params -", err)
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 校验参数
	_, err := logic.Login(l)
	if err != nil {
		if errors.Is(err, logic.ErrorUserNotExist) || errors.Is(err, logic.ErrorWrongPassword) {
			ResponseError(c, CodeWrongUser)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

func EditPasswordHandler(c *gin.Context) {
	// 获取参数
	var Param struct {
		Username      string `json:"username"`
		Password      string `json:"password"`
		NewPassword   string `json:"newPassword"`
		ReNewPassword string `json:"reNewPassword"`
	}
	if err := c.ShouldBindJSON(&Param); err != nil {
		fmt.Println("ShouldBindJSON error -", err)
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 逻辑处理
	if err := logic.EditPassword(Param.Username, Param.Password, Param.NewPassword, Param.ReNewPassword); err != nil {
		if errors.Is(err, logic.ErrorUserNotExist) || errors.Is(err, logic.ErrorWrongPassword) {
			ResponseError(c, CodeWrongUser)
			return
		}
		if errors.Is(err, logic.ErrorSamePassword) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "新旧密码不能相同",
			})
			return
		}
		if errors.Is(err, logic.ErrorTwoPasswordNotSame) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "两次密码输入不相同",
			})
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}
