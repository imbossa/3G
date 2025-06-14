package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imbossa/3G/pkg/logger"
)

// 构建日志信息
func buildRequestMessage(ctx *gin.Context, status int, bodySize int) string {
	var result strings.Builder

	result.WriteString(ctx.ClientIP())
	result.WriteString(" - ")
	result.WriteString(ctx.Request.Method)
	result.WriteString(" ")
	result.WriteString(ctx.Request.RequestURI)
	result.WriteString(" - ")
	result.WriteString(strconv.Itoa(status))
	result.WriteString(" ")
	result.WriteString(strconv.Itoa(bodySize))

	return result.String()
}

// Gin Logger 中间件
func Logger(l logger.Interface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 处理请求
		ctx.Next()

		// 记录响应信息
		status := ctx.Writer.Status()
		bodySize := ctx.Writer.Size() // 响应体字节数

		msg := buildRequestMessage(ctx, status, bodySize)
		l.Info(msg)
	}
}
