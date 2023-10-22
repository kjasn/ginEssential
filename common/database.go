package common

import (
	"Kjasn/ginEssential/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var mysqlLogger logger.Interface

func InitDB() {
	// 连接数据库
	//drivername := "mysql"
	username := "root"
	password := "8520"
	host := "localhost"
	port := "3306"
	database := "gin_essential"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, database)

	// dsn := "root:8520@tcp(127.0.0.1:3306)/ginEssential?charset=utf8mb4&parseTime=True&loc=L=True&loc=Local"

	mysqlLogger = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})

	DB = db

	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	fmt.Println("数据库连接成功")
	db.AutoMigrate(&model.User{})

}
