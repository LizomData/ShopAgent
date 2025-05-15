package router

import (
	"ShopAgent/controller/userController"
	"ShopAgent/util"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(util.Cors())
	group := r.Group("/api/v1/account")
	group.Use()
	{
		group.POST("/login", userController.Login)
		group.POST("/register", userController.Register)
	}

	return r
}
