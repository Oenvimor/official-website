package controller

import (
	"cqupt_hub/logic"
	"cqupt_hub/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UploadHandler(c *gin.Context) {
	// 获取上传文件列表
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("get file failed -", err)
		ResponseError(c, CodeGetImageFail)
		return
	}
	files := form.File["image"] // 获取所有文件
	// 检查文件数量
	if len(files) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "未选择文件",
		})
		return
	}
	// 获取配置参数
	cfg := setting.Conf.QiniuConfig
	accessKey := cfg.AccessKey
	secretKey := cfg.SecretKey
	bucket := cfg.Bucket
	qiniuServer := cfg.QiniuServer
	// 逻辑处理
	uploadResults, err := logic.Upload(c, accessKey, secretKey, bucket, qiniuServer, files)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, uploadResults)
}

func GetImageHandler(c *gin.Context) {
	// 获取图片记录中全部图片地址
	ImageRecords, err := logic.GetImage()
	if err != nil {
		ResponseError(c, CodeGetImageFail)
		return
	}
	ResponseSuccess(c, ImageRecords)
}

func DeleteImageHandler(c *gin.Context) {
	// 获取参数
	id := c.Param("ID")
	Id, _ := strconv.Atoi(id)
	// 逻辑处理
	//// 获取配置参数
	//cfg := setting.Conf.QiniuConfig
	//accessKey := cfg.AccessKey
	//secretKey := cfg.SecretKey
	//bucket := cfg.Bucket
	//if err := logic.QiniuDelete(request.Filename, accessKey, secretKey, bucket); err != nil {
	//	ResponseWithMsg(c, CodeDeleteFailed, err.Error())
	//	return
	//}
	if err := logic.Delete(Id); err != nil {
		ResponseError(c, CodeDeleteFailed)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

func SetHomePageHandler(c *gin.Context) {
	// 获取参数
	var request struct {
		FileName []string `json:"file_names"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		ResponseError(c, CodeGetImageFail)
		return
	}
	// 逻辑处理
	err := logic.SetHomePage(request.FileName)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": err,
		})
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}
