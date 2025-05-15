package controller

import (
	"cqupt_hub/logic"
	"cqupt_hub/models"
	"cqupt_hub/setting"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

func GetDepartmentHandler(c *gin.Context) {
	// 逻辑处理
	data, err := logic.GetDepartment()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

func AddDepartmentHandler(c *gin.Context) {
	// 获取参数
	Username := c.PostForm("personName")
	Email := c.PostForm("email")
	SectionName := c.PostForm("departName")
	Introduction := c.PostForm("personIntro")
	form, err := c.MultipartForm()
	if err != nil {
		ResponseError(c, CodeGetImageFail)
		return
	}
	avatar := form.File["image"] // 获取所有文件
	//// 检查文件数量
	//if len(avatar) == 0 {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "未选择文件",
	//	})
	//	return
	//}
	// 获取配置参数
	cfg := setting.Conf.QiniuConfig
	accessKey := cfg.AccessKey
	secretKey := cfg.SecretKey
	bucket := cfg.Bucket
	qiniuServer := cfg.QiniuServer
	// 上传头像
	var uploadResults []map[string]string
	var avatarURL string
	if avatar != nil {
		uploadResults = logic.QiniuUpload(c, accessKey, secretKey, bucket, qiniuServer, avatar)
		avatarURL = uploadResults[0]["url"]
	}
	// 校验参数
	m := &models.Department{
		Username:     Username,
		Email:        Email,
		Avatar:       avatarURL,
		SectionName:  SectionName,
		Introduction: Introduction,
	}
	// 手动校验
	if m.Username == "" || m.SectionName == "" || m.Email == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "必填字段不能为空",
		})
		return
	}
	// 逻辑处理
	err = logic.AddDepartment(m)
	if err != nil {
		if errors.Is(err, logic.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, m)
}

func EditDepartmentHandler(c *gin.Context) {
	// 获取参数
	ID := c.Param("id")
	UserID, err := strconv.Atoi(ID)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 获取文本字段
	SectionName := c.DefaultPostForm("departName", "")
	Username := c.DefaultPostForm("personName", "")
	Email := c.DefaultPostForm("email", "")
	Introduction := c.DefaultPostForm("personIntro", "")
	// 获取配置参数
	cfg := setting.Conf.QiniuConfig
	accessKey := cfg.AccessKey
	secretKey := cfg.SecretKey
	bucket := cfg.Bucket
	qiniuServer := cfg.QiniuServer
	// 获取图片文件
	form, err := c.MultipartForm()
	if err != nil {
		ResponseError(c, CodeGetImageFail)
		return
	}
	Avatar := form.File["image"]
	UpdateFields := make(map[string]interface{})
	// 如果获取到图片文件才上传头像
	var uploadResults []map[string]string
	var avatarURL string
	if len(Avatar) > 0 {
		// 获取旧头像的url
		var url string
		url, err = logic.GetAvatarUrlByID(UserID)
		// 如果url为空
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				Url := strings.Split(url, "/")
				err = logic.QiniuDelete(Url[1], accessKey, secretKey, bucket)
				if err != nil {
					fmt.Println(err)
					c.JSON(http.StatusOK, gin.H{
						"msg": "删除图片失败",
					})
					return
				}
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg": "获取文件失败",
				})
				return
			}
		}
		// 上传新头像
		uploadResults = logic.QiniuUpload(c, accessKey, secretKey, bucket, qiniuServer, Avatar)
		avatarURL = uploadResults[0]["url"]
		UpdateFields["avatar"] = avatarURL
	}
	// 校验参数
	// 若字段不为空才填充
	if SectionName != "" {
		UpdateFields["section_name"] = SectionName
	}
	if Username != "" {
		UpdateFields["username"] = Username
	}
	if Email != "" {
		UpdateFields["email"] = Email
	}
	if Introduction != "" {
		UpdateFields["introduction"] = Introduction
	}
	// 逻辑处理
	err = logic.EditDepartment(UserID, UpdateFields)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, UpdateFields)
}

func DeleteDepartmentHandler(c *gin.Context) {
	// 获取参数
	ID := c.Param("id")
	UserID, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println("str convert to int failed, err -", err)
		return
	}
	// 获取配置信息
	cfg := setting.Conf.QiniuConfig
	accessKey := cfg.AccessKey
	secretKey := cfg.SecretKey
	bucket := cfg.Bucket
	// 逻辑处理
	// 删除部门的同时删除部门对应的岗位
	// 删除部门
	if err = logic.DeleteDepartment(UserID, accessKey, secretKey, bucket); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "删除图片失败",
		})
		return
	}
	// 获取部门的名称用于删除岗位
	var Name string
	if Name, err = logic.GetSectionNameByID(UserID); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "获取部门名称失败",
		})
		return
	}
	// 删除岗位
	if err = logic.DeletePositionBySectionName(Name); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "删除岗位失败",
		})
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}
