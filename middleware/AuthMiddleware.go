package middleware

import (
	"Kjasn/ginEssential/common"
	"Kjasn/ginEssential/model"
	"Kjasn/ginEssential/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Response(c, http.StatusUnauthorized, 401, "请求未携带token", nil)
			c.Abort()
			return
		}

		// 去掉前缀
		tokenString = tokenString[7:]

		// 解析出 token
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(c, http.StatusUnauthorized, 401, "未授权", nil)
			c.Abort()
			return
		}
		// 通过验证后  获取claims中的userId
		userId := claims.UserId
		DB := common.DB
		var user model.User
		DB.First(&user, userId)
		// 用户不存在
		if user.ID == 0 {
			response.Response(c, http.StatusUnauthorized, 401, "权限不足，用户不存在", nil)
			c.Abort()
			return
		}
		// 用户存在 将user信息写入上下文
		c.Set("user", user)
	}
}
