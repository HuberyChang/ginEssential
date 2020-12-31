package main

import (
	"fmt"
	"ginEssential/common"
	"os"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func main() {
	InitConfig()
	common.InitDB()
	r := gin.Default()
	r = CollectRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	fmt.Println("++++++++++++++++++", workDir+"\\config")
	viper.AddConfigPath(workDir + "\\config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置文件失败")
	}
}
