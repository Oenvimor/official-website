package controller

import (
	"cqupt_hub/logic"
	"cqupt_hub/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
)

// 校验邮箱格式
func isValidEmail(email string) bool {
	// 电子邮件正则表达式
	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// 校验手机号码格式（假设是中国大陆手机号码）
func isValidPhone(phone string) bool {
	// 手机号正则表达式（仅适用于中国大陆 11 位手机号，以 1 开头）
	var phoneRegex = `^1[3-9]\d{9}$`
	re := regexp.MustCompile(phoneRegex)
	return re.MatchString(phone)
}

func AddContactHandler(c *gin.Context) {
	// 获取参数
	contact := new(models.Contact)
	if err := c.ShouldBindJSON(contact); err != nil {
		fmt.Println("shouldBindJson err -", err)
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 手动检验参数
	if isValid := isValidEmail(contact.Email); isValid != true {
		ResponseError(c, CodeInvalidParams)
		return
	}
	if isValid := isValidPhone(contact.PhoneNum); isValid != true {
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 逻辑处理
	if err := logic.AddContact(contact); err != nil {
		fmt.Println("logic.AddContact err -", err)
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

func GetContactHandler(c *gin.Context) {
	// 逻辑处理
	data, err := logic.GetContact()
	if err != nil {
		fmt.Println("logic.GetContact err -", err)
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

func EditContactHandler(c *gin.Context) {
	// 获取参数
	//id := c.Param("id")
	//ID, err := strconv.Atoi(id)
	//if err != nil {
	//	fmt.Println("id convert to int err -", err)
	//	ResponseError(c, CodeServerBusy)
	//	return
	//}
	contact := new(models.Contact)
	if err := c.ShouldBindJSON(contact); err != nil {
		fmt.Println("shouldBindJson err -", err)
		ResponseError(c, CodeInvalidParams)
		return
	}
	UpdateFields := make(map[string]interface{})
	if contact.StudioAddress != "" {
		UpdateFields["studio_address"] = contact.StudioAddress
	}
	if contact.Email != "" {
		UpdateFields["email"] = contact.Email
	}
	if contact.PhoneNum != "" {
		UpdateFields["phone_num"] = contact.PhoneNum
	}
	if contact.QQ != "" {
		UpdateFields["qq"] = contact.QQ
	}
	// 逻辑处理
	if err := logic.EditContact(UpdateFields); err != nil {
		fmt.Println("logic.EditContact err -", err)
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}
