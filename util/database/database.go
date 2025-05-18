package database

import (
	"ShopAgent/model"
	"ShopAgent/model/po"
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var GormDB = initDb()

func initDb() *gorm.DB {
	var configObj = new(model.Config)
	err := ini.MapTo(configObj, "conf/config_db.ini")
	if err != nil {
		log.Fatalf("config failed err: %v", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configObj.User,
		configObj.Password,
		configObj.Host,
		configObj.Port,
		configObj.Db,
	)
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&model.CommodityInfo{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}


	// 自动迁移商品进货出货的表结构
	err = db.AutoMigrate(&po.PurchaseInbound{})
	err = db.AutoMigrate(&po.PurchaseReturn{})

	// 自动迁移表结构
	err = db.AutoMigrate(&model.Supplier{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&model.CommodityCategory{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	
	// 自动迁移表结构
	err = db.AutoMigrate(&model.SalesPerson{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}


	return db

}
