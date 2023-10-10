package controller

import (
	"Kjasn/ginEssential/common"
	"Kjasn/ginEssential/model"
	"Kjasn/ginEssential/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func existTelephone(telephone string) bool {
	var user model.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Register(ctx *gin.Context) {
	//ctx.HTML(http.StatusOK, "register.html", nil)  // why???

	// 获取数据
	// 要求  昵称 手机号 密码
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	pwd := ctx.PostForm("pwd")

	// 数据验证
	if len(name) == 0 { // 为获取到昵称
		name = util.GetRandomString(10) // randString
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
	common.DB.Create(&model.User{
		Name:      name,
		Telephone: telephone,
		Pwd:       pwd,
	})

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}
