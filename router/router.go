package router

import (
	"ShopAgent/controller/commodityCategoryController"
	"ShopAgent/controller/commodityController"
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

	return r
}
