package controller

import (
	"Kjasn/ginEssential/common"
	"Kjasn/ginEssential/dto"
	"Kjasn/ginEssential/model"
	"Kjasn/ginEssential/response"
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
	//name := ctx.PostForm("name")
	//telephone := ctx.PostForm("telephone")
	//pwd := ctx.PostForm("pwd")
	var requestUser = model.User{}
	err := ctx.Bind(&requestUser)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, "数据验证错误", nil)
		return
	}
	name := requestUser.Name
	telephone := requestUser.Telephone
	pwd := requestUser.Pwd

	// 数据验证
	if len(name) == 0 { // 为获取到昵称
		name = util.GetRandomString(10) // randString
	}

	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, "手机号必须为11位", nil)
		return // 终止
	}

	if len(pwd) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, "密码至少为6位", nil)
		return // 终止
	}

	// 检验 手机号是否以及存在
	if existTelephone(telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, "手机号已注册", nil)
		return // 终止
	}

	fmt.Printf(name, telephone, pwd)

	// 通过校验后 创建用户

	// 加密密码
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, "加密失败", nil)
		return // 终止
	}

	common.DB.Create(&model.User{
		Name:      name,
		Telephone: telephone,
		Pwd:       string(hashPwd), // 保存加密后的密码
	})

	if common.DB.Error != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, "注册失败", nil)
		fmt.Println(common.DB.Error)
		return // 终止
	}

	// 返回结果
	response.Success(ctx, "注册成功", nil)
}

func Login(ctx *gin.Context) {
	// 通过 手机号 + 密码 登录
	telephone := ctx.PostForm("telephone")
	pwd := ctx.PostForm("pwd")

	// 数据验证

	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, "手机号必须为11位", nil)
		return // 终止
	}

	if len(pwd) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, "密码至少为6位", nil)
		return // 终止
	}

	var user model.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, "该手机号码未注册", nil)
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(pwd)); err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, "密码错误", nil)
		return
	}

	// 生成token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, "系统错误", nil)
		fmt.Println("生成token失败", err) // 记录日志
		return
	}

	// 返回结果
	response.Success(ctx, "登录成功", gin.H{"token": token})
}

func Info(ctx *gin.Context) {
	// 用户已通过验证	直接返回用户信息 只显示 name 和 telephone
	user, _ := ctx.Get("user")
	response.Success(ctx, "获取成功", gin.H{"user": dto.ToUserDto(user.(model.User))})
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"data": gin.H{"user": dto.ToUserDto(user.(model.User))}, // 断言
	//})
}
