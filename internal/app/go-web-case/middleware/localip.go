package middleware

//
//import (
//	"github.com/gin-gonic/gin"
//	"net"
//	"net/http"
//)
//
//// 仅仅限制内网访问
//func OnlyLocalIP() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		clientIP := utility.GetClientIP(c.Request)
//		if net.IsLan(clientIP) {
//			c.Next()
//		} else {
//			c.AbortWithStatus(http.StatusForbidden)
//			return
//		}
//	}
//}
