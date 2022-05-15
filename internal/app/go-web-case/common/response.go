package common

// JSONResponse Response对象结构
type JSONResponse struct {
	Errno     string      `json:"errno"`
	ErrMsg    string      `json:"errmsg"`
	RequestID string      `json:"request_id"`
	Data      interface{} `json:"data"`
}

// GetSuccess 获取成功的Response信息
// @param data interface{} response数据
// @param  requestID string 请求ID
// @param  msg string response消息
// @return    	*JSONResponse   response对象
func GetSuccess(data interface{}, requestID string, msg string) *JSONResponse {
	return &JSONResponse{
		Errno:     Success,
		ErrMsg:    msg,
		RequestID: requestID,
		Data:      data,
	}
}

// GetError 获取错误的Response信息
// @param  errorNo string 错误码
// @param data interface{} response数据
// @param  requestID string 请求ID
// @return    	*JSONResponse   response对象
func GetError(errorNo string, data interface{}, requestID string, args ...string) *JSONResponse {
	return &JSONResponse{
		Errno:     errorNo,
		ErrMsg:    GetErrorString(errorNo, args...),
		RequestID: requestID,
		Data:      data,
	}
}
