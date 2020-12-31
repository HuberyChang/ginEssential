/*
* @Author: HuberyChang
* @Date: 2020/12/31 16:01
 */

package middleware

import (
	"fmt"
	"ginEssential/response"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(ctx, nil, fmt.Sprint(err))
			}
		}()
		ctx.Next()
	}
}
