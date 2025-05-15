package controller

import (
	"cqupt_hub/logic"
	"cqupt_hub/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func GetPositionHandler(c *gin.Context) {
	// 逻辑处理
	data, err := logic.GetPosition()
	if err != nil {
		ResponseError(c, CodeGetPositionFail)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

func AddPositionHandler(c *gin.Context) {
	// 获取参数
	p := new(models.Position)
	if err := c.ShouldBindJSON(p); err != nil {
		fmt.Println("bind json err -", err)
		ResponseError(c, CodeServerBusy)
		return
	}
	// 逻辑处理
	if err := logic.AddPosition(p); err != nil {
		if errors.Is(err, logic.ErrorPositionExist) {
			ResponseError(c, CodePositionExist)
			return
		}
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

func EditPositionHandler(c *gin.Context) {
	// 获取参数
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("id convert to int err -", err)
		return
	}
	p := new(models.Position)
	if err = c.ShouldBindJSON(p); err != nil {
		fmt.Println("bind json err -", err)
		ResponseError(c, CodeServerBusy)
		return
	}
	// 校验参数
	UpdateFields := make(map[string]interface{})
	if p.PositionName != "" {
		UpdateFields["position_name"] = p.PositionName
	}
	if p.Requirement != "" {
		UpdateFields["requirement"] = p.Requirement
	}
	if p.Delivery != "" {
		UpdateFields["delivery"] = p.Delivery
	}
	// 逻辑处理
	if err = logic.EditPosition(ID, UpdateFields); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "修改部门失败",
		})
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

func DeletePositionHandler(c *gin.Context) {
	// 获取参数
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("id convert to int err -", err)
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 逻辑处理
	if err = logic.DeletePosition(ID); err != nil {
		fmt.Println("delete position err -", err)
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}
