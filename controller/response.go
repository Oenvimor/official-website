package controller

import "github.com/gin-gonic/gin"

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(200, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

func ResponseWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(200, &ResponseData{
		Code: code,
		Data: code.Msg(),
		Msg:  msg,
	})
}
