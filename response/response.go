package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus, code int, data gin.H, msg string) {
	// code是业务返回状态码，httpStatus是http的状态码
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
