package router

import (
	"ShopAgent/controller/commodityController"
	"ShopAgent/controller/purchase_controller"
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

	group = r.Group("/api/v1/purchase")
	group.Use(util.AuthMiddleware())
	{
		// 进货管理
		group.POST("/inbound", purchase_controller.CreateInbound)
		group.GET("/inbound/list", purchase_controller.GetInboundList)

		// 退货管理
		group.POST("/return", purchase_controller.CreateReturn)
		group.GET("/return/list", purchase_controller.GetReturnList)
	}

	return r
}
