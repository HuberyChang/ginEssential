package middleware

import (
	"ginEssential/common"
	"ginEssential/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足0"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足1"})
			ctx.Abort()
			return
		}

		// 验证通过后获取claims 中的userID
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		if userId == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足2"})
			ctx.Abort()
			return
		}

		// 如果存在，将user的信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
