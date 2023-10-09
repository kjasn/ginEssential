package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `json:"name" gorm:"type:varchar(20);not null"`
	Telephone string `json:"telephone" gorm:"type:varchar(11);not null;unique"`
	Pwd       string `json:"pwd" gorm:"type:varchar(20);not null"`
}

func main() {

	r := gin.Default()

	r.POST("/api/auth/register", func(ctx *gin.Context) {
		// 获取数据
		// 要求  昵称 手机号 密码
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		pwd := ctx.PostForm("pwd")

		// 数据验证
		if len(name) == 0 { // 为获取到昵称
			name = getRandomString(10)
			fmt.Println(name)
		}

		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity,
				gin.H{"msg": "手机号必须为11位"})
			return // 终止
		}

		if len(pwd) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "密码至少为6位"})
			return // 终止
		}

		// 检验 手机号是否以及存在
		if existTelephone(telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "手机号已经存在"})
			return // 终止
		}

		// 通过校验后 创建用户
		DB.Debug().Create(&User{
			Name:      name,
			Telephone: telephone,
			Pwd:       pwd,
		})

		// 返回结果
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})

	err := r.Run(":80")
	if err != nil {
		fmt.Println(err)
	}
}

func getRandomString(sz int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*")
	var ret = make([]byte, sz)
	rand.Seed(time.Now().Unix())
	for i := range ret {
		ret[i] = letters[rand.Intn(len(letters))]
	}

	return string(ret)
}

func existTelephone(telephone string) bool {
	var user User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
