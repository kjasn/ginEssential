package common

import (
	"Kjasn/ginEssential/model"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var mysqlLogger logger.Interface

func InitDB() {
	// 连接数据库		读取config文件获取配置信息
	//drivername := "mysql"
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		username, password, host, port, database, charset)
	// fmt.Println(dsn)

	//dsn := "root:8520@tcp(127.0.0.1:3306)/gin_essential?charset=utf8mb4&parseTime=True"

	mysqlLogger = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})

	DB = db

	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	fmt.Println("数据库连接成功")
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return
	}

}
