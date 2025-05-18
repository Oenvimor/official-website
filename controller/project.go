package controller

import (
	"cqupt_hub/logic"
	"cqupt_hub/models"
	"cqupt_hub/setting"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func GetProjectHandler(c *gin.Context) {
	// 逻辑处理
	data, err := logic.GetProject()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

func AddProjectHandler(c *gin.Context) {
	// 获取参数
	Name := c.PostForm("projectName")
	Path := c.PostForm("projectPath")
	form, err := c.MultipartForm()
	if err != nil {
		ResponseError(c, CodeGetImageFail)
		return
	}
	projectImage := form.File["image"] // 获取所有文件
	//// 检查文件数量
	//if len(projectImage) == 0 {
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
	// 上传项目展示图片
	var uploadResults []map[string]string
	var ImageUrl string
	if projectImage != nil {
		uploadResults = logic.QiniuUpload(c, accessKey, secretKey, bucket, qiniuServer, projectImage)
		for _, item := range uploadResults {
			if item["status"] == "失败" {
				fmt.Printf("err:%v", item["error"])
				ResponseError(c, CodeUploadFail)
				return
			}
		}
		ImageUrl = uploadResults[0]["url"]
	}
	// 校验参数
	m := &models.Project{
		ProjectName:  Name,
		ProjectPath:  Path,
		ProjectImage: ImageUrl,
	}
	// 手动校验
	if m.ProjectName == "" || m.ProjectPath == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "必填字段不能为空",
		})
		return
	}
	// 逻辑处理
	err = logic.AddProject(m)
	if err != nil {
		if errors.Is(err, logic.ErrorProjectExist) {
			ResponseError(c, CodeProjectExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

func EditProjectHandler(c *gin.Context) {
	// 获取参数
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 获取文本字段
	projectName := c.DefaultPostForm("projectName", "")
	projectPath := c.DefaultPostForm("projectPath", "")
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
	Image := form.File["image"]
	UpdateFields := make(map[string]interface{})
	// 如果获取到图片文件才上传
	var uploadResults []map[string]string
	var ProjectURL string
	if len(Image) > 0 {
		// 删除旧图片
		// 获取项目的旧图片
		var url string
		url, err = logic.GetImageUrlByID(ID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "获取旧头像失败",
			})
			return
		}
		Url := strings.Split(url, "/")
		// 从七牛云删除
		err = logic.QiniuDelete(Url[1], accessKey, secretKey, bucket)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"msg": "删除图片失败",
			})
			return
		}
		// 上传新图片
		uploadResults = logic.QiniuUpload(c, accessKey, secretKey, bucket, qiniuServer, Image)
		ProjectURL = uploadResults[0]["url"]
		UpdateFields["project_image"] = ProjectURL
	}
	// 校验参数
	// 若字段不为空才填充
	if projectName != "" {
		UpdateFields["project_name"] = projectName
	}
	if projectPath != "" {
		UpdateFields["project_path"] = projectPath
	}
	// 逻辑处理
	err = logic.EditProject(ID, UpdateFields)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, UpdateFields)
}

func DeleteProjectHandler(c *gin.Context) {
	// 获取参数
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
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
	if err = logic.DeleteProject(ID, accessKey, secretKey, bucket); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "删除图片失败",
		})
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

func DisplayProjectHandler(c *gin.Context) {
	// 获取参数
	var request struct {
		ProjectID []int `json:"projectID"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		ResponseError(c, CodeGetImageFail)
		return
	}
	// 逻辑处理
	err := logic.DisplayProject(request.ProjectID)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "展示项目失败",
		})
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}
