package controller

import (
	"cqupt_hub/logic"
	"cqupt_hub/models"
	"cqupt_hub/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
)

func GetGameHandler(c *gin.Context) {
	// 逻辑处理
	data, err := logic.GetGame()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

func AddGameHandler(c *gin.Context) {
	// 获取参数
	gameName := c.PostForm("gameName")
	gamePath := c.PostForm("gamePath")
	form, err := c.MultipartForm()
	if err != nil {
		ResponseError(c, CodeGetImageFail)
		return
	}
	gameImage := form.File["image"] // 获取所有文件
	//// 检查文件数量
	//if len(gameImage) == 0 {
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
	var ImageUrl string
	var uploadResults []map[string]string
	if gameImage != nil {
		uploadResults = logic.QiniuUpload(c, accessKey, secretKey, bucket, qiniuServer, gameImage)
		ImageUrl = uploadResults[0]["url"]
	}

	// 校验参数
	m := &models.Game{
		GameName:  gameName,
		GamePath:  gamePath,
		GameImage: ImageUrl,
	}
	// 手动校验
	if m.GameName == "" || m.GamePath == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "必填字段不能为空",
		})
		return
	}
	// 逻辑处理
	err = logic.AddGame(m)
	if err != nil {
		if errors.Is(err, logic.ErrorGameExist) {
			ResponseError(c, CodeGameExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, m)
}

func EditGameHandler(c *gin.Context) {
	// 获取参数
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 获取文本字段
	gameName := c.DefaultPostForm("gameName", "")
	gamePath := c.DefaultPostForm("gamePath", "")
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
	var gameURL string
	if len(Image) > 0 {
		// 删除旧图片
		// 获取项目的旧图片
		var url string
		url, err = logic.GetImageUrl(ID)
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
		// 上传新头像
		uploadResults = logic.QiniuUpload(c, accessKey, secretKey, bucket, qiniuServer, Image)
		gameURL = uploadResults[0]["url"]
		UpdateFields["game_image"] = gameURL
	}
	// 校验参数
	// 若字段不为空才填充
	if gameName != "" {
		UpdateFields["game_name"] = gameName
	}
	if gamePath != "" {
		UpdateFields["game_path"] = gamePath
	}
	// 逻辑处理
	err = logic.EditGame(ID, UpdateFields)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, UpdateFields)
}

func DeleteGameHandler(c *gin.Context) {
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
	if err = logic.DeleteGame(ID, accessKey, secretKey, bucket); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "删除图片失败",
		})
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

func DisplayGameHandler(c *gin.Context) {
	// 获取参数
	var request struct {
		GameID []int `json:"gameID"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		ResponseError(c, CodeGetImageFail)
		return
	}
	// 逻辑处理
	err := logic.DisplayGame(request.GameID)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "展示项目失败",
		})
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}
