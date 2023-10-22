package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
格式：
{
    "code": xxx
    "msg": xxx
    "data": xxx
}
*/

// Response  封装  自定义 http返回格式
func Response(ctx *gin.Context, httpStatus int, code int, msg string, data interface{}) {
	ctx.JSON(httpStatus, gin.H{"code": code, "msg": msg, "data": data})
}

//------------------ 几个常用的返回 ---------------------

// Success 返回成功
func Success(ctx *gin.Context, msg string, data interface{}) {
	Response(ctx, http.StatusOK, 200, msg, data)
}

// Fail 返回失败
func Fail(ctx *gin.Context, msg string, data interface{}) {
	Response(ctx, http.StatusOK, 400, msg, data) // http状态仍然是200	另外返回失败的状态码400
}
