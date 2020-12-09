package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 422, "msg": "权限不足"})
			ctx.Abort()
			return
		}

	}
}
