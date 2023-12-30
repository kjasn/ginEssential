package main

import (
	"Kjasn/ginEssential/common"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig() // 读取配置文件
	fmt.Println("reading config file successfully")
	common.InitDB() // 初始化数据库
	fmt.Println("init db successfully")

	r := gin.Default()

	r = CollectRouter(r)

	port := viper.GetString("server.port")
	fmt.Println(port)
	err := r.Run(":" + port)
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
	viper.SetConfigType("yaml")              // 配置文件类型
	viper.AddConfigPath(workDir + "/config") // 配置文件路径
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
