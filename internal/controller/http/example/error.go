package http

import (
	"github.com/gin-gonic/gin"
	"github.com/imbossa/3G/internal/controller/http/example/response"
)

func errorResponse(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, response.Error{Error: msg})
}
