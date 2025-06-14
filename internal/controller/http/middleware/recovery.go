package middleware

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imbossa/3G/pkg/logger"
)

func buildPanicMessage(ctx *gin.Context, err interface{}) string {
	var result strings.Builder

	result.WriteString(ctx.ClientIP())
	result.WriteString(" - ")
	result.WriteString(ctx.Request.Method)
	result.WriteString(" ")
	result.WriteString(ctx.Request.RequestURI)
	result.WriteString(" PANIC DETECTED: ")
	result.WriteString(fmt.Sprintf("%v\n%s\n", err, debug.Stack()))

	return result.String()
}

// Gin Recovery 中间件
func Recovery(l logger.Interface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录 panic 信息和堆栈
				l.Error(buildPanicMessage(ctx, err))
				// 返回 500 状态
				ctx.AbortWithStatus(500)
			}
		}()
		ctx.Next()
	}
}
