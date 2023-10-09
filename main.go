package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func getRandomString(sz int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*")
	var ret = make([]byte, sz)
	rand.Seed(time.Now().Unix())
	for i := range ret {
		ret[i] = letters[rand.Intn(len(letters))]
	}

	return string(ret)
}

func main() {
	r := gin.Default()

	r.GET("/api/auth/register", func(ctx *gin.Context) {
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

//func main() {
//	a := rand.Intn(100)
//	fmt.Println(a)
//}
