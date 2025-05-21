package router

import (
	"ShopAgent/controller/commodityCategoryController"
	"ShopAgent/controller/commodityController"
	"ShopAgent/controller/imageUploadController"
	"ShopAgent/controller/purchase_controller"
	"ShopAgent/controller/salespersonController"
	"ShopAgent/controller/supplierController"
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
		group.PUT("/inbound/update", purchase_controller.UpdateInbound)
		group.DELETE("/inbound/delete", purchase_controller.DeleteInbound)

		// 退货管理
		group.POST("/return", purchase_controller.CreateReturn)
		group.GET("/return/list", purchase_controller.GetReturnList)
		group.PUT("/return/update", purchase_controller.UpdateReturn)
		group.DELETE("/return/delete", purchase_controller.DeleteReturn)
	}
	group = r.Group("/api/v1/supplier")
	group.Use(util.AuthMiddleware())
	{
		group.POST("/create", supplierController.Create)
		group.POST("/query", supplierController.Query)
		group.POST("/update", supplierController.Update)
		group.POST("/delete", supplierController.Delete)
	}
	group = r.Group("/api/v1/commodityCategory")
	group.Use(util.AuthMiddleware())
	{
		group.POST("/create", commodityCategoryController.Create)
		group.POST("/query", commodityCategoryController.Query)
		group.POST("/update", commodityCategoryController.Update)
		group.POST("/delete", commodityCategoryController.Delete)
	}

	group = r.Group("/api/v1/salesperson")
	group.Use(util.AuthMiddleware())
	{
		group.POST("/create", salespersonController.Create)
		group.POST("/query", salespersonController.Query)
		group.POST("/update", salespersonController.Update)
		group.POST("/delete", salespersonController.Delete)
	}

	// 暴露静态文件目录
	r.Static(imageUploadController.Uploader.Config.PublicPath, imageUploadController.Uploader.Config.UploadDir)
	group = r.Group("/api/v1/image")
	group.Use(util.AuthMiddleware())
	{
		group.POST("/upload", imageUploadController.Upload)
		group.POST("/query", imageUploadController.Query)
	}

	return r
}
