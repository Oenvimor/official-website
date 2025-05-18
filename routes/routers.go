package routes

import (
	"cqupt_hub/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	// 后台管理
	r.POST("/register", controller.RegisterHandler)   // 管理员注册
	r.POST("/login", controller.LoginHandler)         // 管理员登录
	r.PUT("/account", controller.EditPasswordHandler) // 修改密码

	v1 := r.Group("/image")                                 // 首页设置
	v1.POST("/uploads", controller.UploadHandler)           // 图片上传到图片库
	v1.DELETE("/delete/:id", controller.DeleteImageHandler) // 从图片库删除图片
	v1.GET("/acquire", controller.GetImageHandler)          // 获取图片库所有图片
	v1.PUT("/homepage", controller.SetHomePageHandler)      // 设置首页图片

	v2 := r.Group("/department")                                 // 部门管理
	v2.GET("/acquire", controller.GetDepartmentHandler)          // 获取部门信息
	v2.POST("/append", controller.AddDepartmentHandler)          // 添加部门信息
	v2.PUT("/edit/:id", controller.EditDepartmentHandler)        // 编辑部门信息
	v2.DELETE("/delete/:id", controller.DeleteDepartmentHandler) // 删除部门信息

	v3 := r.Group("/position")                                 // 岗位管理
	v3.GET("/acquire", controller.GetPositionHandler)          // 获取岗位信息
	v3.POST("/append", controller.AddPositionHandler)          // 添加岗位信息
	v3.PUT("/edit/:id", controller.EditPositionHandler)        // 编辑岗位信息
	v3.DELETE("/delete/:id", controller.DeletePositionHandler) // 删除岗位信息

	v4 := r.Group("/project")                                 // 项目管理
	v4.GET("/acquire", controller.GetProjectHandler)          // 获取项目信息
	v4.POST("/append", controller.AddProjectHandler)          // 添加项目信息
	v4.PUT("/edit/:id", controller.EditProjectHandler)        // 编辑项目信息
	v4.DELETE("/delete/:id", controller.DeleteProjectHandler) // 删除项目信息
	v4.PUT("/display", controller.DisplayProjectHandler)      // 展示项目信息

	v5 := r.Group("/game")                                 // 游戏管理
	v5.GET("/acquire", controller.GetGameHandler)          // 获取游戏信息
	v5.POST("/append", controller.AddGameHandler)          // 添加游戏信息
	v5.PUT("/edit/:id", controller.EditGameHandler)        // 编辑游戏信息
	v5.DELETE("/delete/:id", controller.DeleteGameHandler) // 删除游戏信息
	v5.PUT("/display", controller.DisplayGameHandler)      // 展示游戏信息

	v6 := r.Group("/contact")                        // 联系方式管理
	v6.POST("/append", controller.AddContactHandler) // 添加联系方式
	v6.GET("/acquire", controller.GetContactHandler) // 获取联系方式
	v6.PUT("/edit", controller.EditContactHandler)   // 修改联系方式

	r.Run(fmt.Sprintf("%s:%d", viper.GetString("app.host"), viper.GetInt("app.port")))
	return r
}
