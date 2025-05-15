package logic

import (
	"cqupt_hub/dao/mysql"
	"cqupt_hub/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"path/filepath"
	"strings"
)

const maxFileSize = 5 << 20

var uploadStatus = "失败"

func Upload(c *gin.Context, accessKey string, secretKey string, bucket string, qiniuServer string, files []*multipart.FileHeader) ([]map[string]string, error) {
	// 获取图片上传结果
	uploadResults := QiniuUpload(c, accessKey, secretKey, bucket, qiniuServer, files)
	// 将图片元数据绑定到结构体
	for n, uploadResult := range uploadResults {
		i := &models.ImageRecord{
			FileName:   uploadResult["filename"],
			URL:        uploadResult["url"],
			IsHomepage: false,
		}
		if uploadResults[n]["status"] != uploadStatus {
			// 将图片数据插入数据库
			err := mysql.Upload(i)
			if err != nil {
				fmt.Println("insert into database failed, err -", err)
			}
		}
	}
	return uploadResults, nil
}

func Delete(filename string) error {
	err := mysql.Delete(filename)
	if err != nil {
		fmt.Println("delete file failed, err -", err)
		return err
	}
	return nil
}
func GetImage() ([]models.ImageRecord, error) {
	ImageRecords, err := mysql.GetImage()
	if err != nil {
		fmt.Println("get image failed, err -", err)
		return nil, err
	}
	return ImageRecords, nil
}

func SetHomePage(fileName []string) error {
	err := mysql.SetHomePage(fileName)
	if err != nil {
		fmt.Println("set homepage failed, err -", err)
		return err
	}
	return nil
}

func QiniuDelete(filename string, accessKey string, secretKey string, bucket string) error {
	// 初始化七牛云 MAC 对象
	mac := qbox.NewMac(accessKey, secretKey)
	// 配置存储空间管理器
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong, // 根据你的存储区域选择
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	// 删除文件
	err := bucketManager.Delete(bucket, filename)
	if err != nil {
		return err
	}
	return nil
}

func QiniuUpload(c *gin.Context, accessKey string, secretKey string, bucket string, qiniuServer string, files []*multipart.FileHeader) []map[string]string {
	mac := qbox.NewMac(accessKey, secretKey) // 创建mac对象用于上传凭证
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	formUploader := storage.NewFormUploader(&cfg) // 创建一个表单上传器的对象
	uploadResults := make([]map[string]string, 0) // 用于存储上传结果
	// 遍历每个文件并上传
	for _, file := range files {
		// 限制不能重复上传
		count, err := mysql.CheckImageByName(file.Filename)
		if err != nil {
			fmt.Println("mysql.CheckImageByName err -", err)
			return nil
		}
		if count > 0 {
			uploadResults = append(uploadResults, map[string]string{
				"filename": file.Filename,
				"status":   "失败",
				"error":    "图片已存在，无法重复上传",
			})
			continue
		}
		// 限制文件大小
		if file.Size > maxFileSize {
			uploadResults = append(uploadResults, map[string]string{
				"filename": file.Filename,
				"status":   "失败",
				"error":    "文件大小不能超过5MB",
			})
			continue
		}
		// 限制文件类型
		allowedType := map[string]bool{
			"image/jpeg": true,
			"image/png":  true,
		}
		if !allowedType[file.Header.Get("Content-Type")] {
			uploadResults = append(uploadResults, map[string]string{
				"filename": file.Filename,
				"status":   "失败",
				"error":    "只允许上传 JPEG 或 PNG 图片",
			})
			continue
		}
		// 打开文件
		fileReader, err := file.Open()
		if err != nil {
			uploadResults = append(uploadResults, map[string]string{
				"filename": file.Filename,
				"status":   "失败",
				"error":    "文件读取失败",
			})
			continue
		}
		// 生成上传凭证
		putPolicy := storage.PutPolicy{
			Scope: bucket,
		}
		upToken := putPolicy.UploadToken(mac)
		// 上传到七牛云
		ret := storage.PutRet{}
		putExtra := storage.PutExtra{}
		key := filepath.Join(file.Filename) // 指定存储路径和文件名
		err = formUploader.Put(c.Request.Context(), &ret, upToken, key, fileReader, file.Size, &putExtra)
		if err != nil {
			uploadResults = append(uploadResults, map[string]string{
				"filename": file.Filename,
				"status":   "失败",
				"error":    fmt.Sprintf("文件上传到七牛云失败: %s", err),
			})
			continue
		}
		URL := fmt.Sprintf("%s/%s", qiniuServer, ret.Key)
		fileURL := strings.Replace(URL, "\\", "%5C", -1)
		uploadResults = append(uploadResults, map[string]string{
			"filename": file.Filename,
			"status":   "成功",
			"url":      fileURL,
		})
	}
	return uploadResults
}
