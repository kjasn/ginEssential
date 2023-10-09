package main

import (
	"Kjasn/ginEssential/controller"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	//r.POST("/api/auth/login", controller.Login)

	return r
}