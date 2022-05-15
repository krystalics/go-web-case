package common

import "fmt"

const (
	Success                = "0"
	InternalServerError    = "InternalServerError"
	IamAuthenticateError   = "IamAuthenticateError"
	IamAuthError           = "IamAuthError"
	AccessForbiddenOnlyLan = "AccessForbiddenOnlyLan"
	RequestParameterError  = "RequestParameterError"
	DatabaseQueryError     = "DatabaseQueryError"
)

var errDefine = map[string]string{
	"InternalServerError":    "服务器内部错误",
	"IamAuthenticateError":   "IAM认证失败",
	"IamAuthError":           "IAM鉴权失败",
	"AccessForbiddenOnlyLan": "只允许内网访问",
	"RequestParameterError":  "请求参数[%s]错误",
	"DatabaseQueryError":     "数据库查询错误",
}

// GetErrorString 根据错误码获取错误信息
// @param errorNo string 错误码
// @return    	string   "错误码对应的错误信息"
func GetErrorString(errorNo string, args ...string) string {
	errmsg, ok := errDefine[errorNo]
	if !ok {
		return "未知错误"
	}
	if len(args) > 0 {
		return fmt.Sprintf(errmsg, args)
	}
	return errmsg
}
