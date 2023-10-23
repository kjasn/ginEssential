package main

import (
	"Kjasn/ginEssential/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()    // 读取配置文件
	common.InitDB() // 初始化数据库

	r := gin.Default()

	r = CollectRouter(r)

	port := viper.GetString("server.port")
	err := r.Run(":%s", port)
	if err != nil {
		fmt.Println(err)
	}
}

func InitConfig() {
	workDir, err := os.Getwd() // 获取当前工作目录
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("application")       // 配置文件名称
	viper.SetConfigType("yml")               // 配置文件类型
	viper.AddConfigPath(workDir + "/config") // 配置文件路径
}
