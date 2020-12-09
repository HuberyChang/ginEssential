package main

import (
	"ginEssential/common"

	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()
	r := gin.Default()
	r = CollectRouter(r)
	panic(r.Run())
}
