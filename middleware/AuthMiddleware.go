package middleware

import (
	"Kjasn/ginEssential/common"
	"Kjasn/ginEssential/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "请求未携带token",
			})
			c.Abort()
			return
		}

		// 去掉前缀
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
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
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		// 用户存在 将user信息写入上下文
		c.Set("user", user)
	}
}
