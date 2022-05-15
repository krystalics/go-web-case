package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// JSONResponse Response对象结构
type JSONResponse struct {
	Errno  string      `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func ResSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, GetSuccess(data, "success"))
}

func ResFailed(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, GetError(GetErrorString(InternalServerError), nil, "failed"))
}

// GetSuccess 获取成功的Response信息
// @param data interface{} response数据
// @param  requestID string 请求ID
// @param  msg string response消息
// @return    	*JSONResponse   response对象
func GetSuccess(data interface{}, msg string) *JSONResponse {
	return &JSONResponse{
		Errno:  Success,
		ErrMsg: msg,
		Data:   data,
	}
}

// GetError 获取错误的Response信息
// @param  errorNo string 错误码
// @param data interface{} response数据
// @param  requestID string 请求ID
// @return    	*JSONResponse   response对象
func GetError(errorNo string, data interface{}, args ...string) *JSONResponse {
	return &JSONResponse{
		Errno:  errorNo,
		ErrMsg: GetErrorString(errorNo, args...),
		Data:   data,
	}
}
