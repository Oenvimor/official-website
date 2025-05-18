package controller

type ResCode int64

const (
	CodeSuccess ResCode = 200 + iota
	CodeInvalidParams
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidAuth
	CodeWrongUser
	CodeUserExist
	CodeUserNotExist
	CodeGetImageFail
	CodeDeleteFailed
	CodePositionExist
	CodePositionNotExist
	CodeGetPositionFail
	CodeProjectExist
	CodeProjectNotExist
	CodeGameExist
	CodeUploadFail
)

var CodeMsgMap = map[ResCode]string{
	CodeSuccess:          "success",
	CodeInvalidParams:    "请求参数有误",
	CodeServerBusy:       "服务繁忙",
	CodeNeedLogin:        "需要登陆",
	CodeInvalidAuth:      "无效的token",
	CodeWrongUser:        "用户名或密码错误",
	CodeUserExist:        "用户已存在",
	CodeUserNotExist:     "用户不存在",
	CodeGetImageFail:     "获取图片失败",
	CodeDeleteFailed:     "删除图片失败",
	CodePositionExist:    "部门已存在",
	CodePositionNotExist: "部门不存在",
	CodeGetPositionFail:  "获取部门失败",
	CodeProjectExist:     "项目已存在",
	CodeProjectNotExist:  "项目不存在",
	CodeGameExist:        "游戏已存在",
	CodeUploadFail:       "上传图片失败",
}

func (code ResCode) Msg() string {
	msg, ok := CodeMsgMap[code]
	if !ok {
		return CodeMsgMap[CodeServerBusy]
	}
	return msg
}
