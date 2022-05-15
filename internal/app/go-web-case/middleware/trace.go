package middleware

//
//import (
//	"github.com/gin-gonic/gin"
//	"bigdata-test/common/trace"
//)
//
//// 处理api请求的Auth验证
//func TraceRequest() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		requestID := c.Request.Header.Get(trace.HeaderKeyRequestID)
//		if requestID == "" {
//			requestID = trace.GenerateRequestID()
//		}
//		//TODO 后续根据公司的trace规范完善其他信息
//		traceInfo := &trace.TraceInfo{
//			TraceID:      requestID,
//			SpanID:       "",
//			ParentSpanID: "",
//		}
//		c.Set(trace.ContextKeyTrace, traceInfo)
//		c.Next()
//	}
//}
