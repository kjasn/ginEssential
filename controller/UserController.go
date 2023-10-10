package controller

import (
	"Kjasn/ginEssential/common"
	"Kjasn/ginEssential/model"
	"Kjasn/ginEssential/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	// 获取数据
	// 要求  昵称 手机号 密码
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	pwd := ctx.PostForm("pwd")

	// 数据验证
	if len(name) == 0 { // 为获取到昵称
		name = util.GetRandomString(10) // randString
	}

	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity,
			gin.H{"code": 422, "msg": "手机号必须为11位"})
		return // 终止
	}

	if len(pwd) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码至少为6位",
		})
		return // 终止
	}

	// 检验 手机号是否以及存在
	if existTelephone(telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "手机号已经存在"})
		return // 终止
	}

	fmt.Printf(name, telephone, pwd)

	// 通过校验后 创建用户

	// 加密密码
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密失败"})
		return // 终止
	}

	common.DB.Create(&model.User{
		Name:      name,
		Telephone: telephone,
		Pwd:       string(hashPwd), // 保存加密后的密码
	})

	if common.DB.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "注册失败"})
		fmt.Println(common.DB.Error)
		return // 终止
	}

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func Login(ctx *gin.Context) {
	// 通过 手机号 + 密码 登录
	telephone := ctx.PostForm("telephone")
	pwd := ctx.PostForm("pwd")

	// 数据验证
	var user model.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该手机号码未注册"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(pwd)); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码错误"})
		return
	}

	// 生成token
	token := "111"

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})
}
