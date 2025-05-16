package router

import (
	"ShopAgent/controller/commodityController"
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

	group = r.Group("/api/v1/commodity")
	group.Use(util.AuthMiddleware())
	{
		group.POST("/create", commodityController.Create)
		group.POST("/query", commodityController.Query)
		group.POST("/update", commodityController.Update)
		group.POST("/delete", commodityController.Delete)
	}

	return r
}
