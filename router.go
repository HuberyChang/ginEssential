package main

import (
	"ginEssential/controller"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	return r
}
