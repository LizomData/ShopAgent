package commodityDbService

import (
	"ShopAgent/model"
	"ShopAgent/util/database"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"
)

var Instance = GetInstance()

func GetInstance() *CommodityDbService {
	// 初始化节点（确保每个服务实例的 nodeID 唯一）
	nodeID := int64(2)
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		fmt.Println("初始化雪花节点失败: " + err.Error())
		return nil
	}
	return &CommodityDbService{node}
}

type CommodityDbService struct {
	node *snowflake.Node
}

func (u *CommodityDbService) Create(commodity_info *model.CommodityInfo) (*model.CommodityInfo, error) {
	commodity_info.ID = u.GenerateSnowflakeID()
	commodity_info.CreatedAt = time.Now().Unix()
	commodity_info.UpdatedAt = time.Now().Unix()
	result := database.GormDB.Create(&commodity_info)
	return commodity_info, result.Error
}
func (s *CommodityDbService) Query(user_id int64, page, size int) ([]model.CommodityInfo, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}

	offset := (page - 1) * size
	var commodity_infos []model.CommodityInfo

	err := database.GormDB.
		Where("user_id = ?", user_id).
		Offset(offset).
		Limit(size).
		Find(&commodity_infos).Error

	return commodity_infos, err
}

func (s *CommodityDbService) GetById(user_id, id int64) (model.CommodityInfo, error) {
	var commodity_info model.CommodityInfo
	err := database.GormDB.
		Where("user_id = ?", user_id).
		First(&commodity_info, "id = ?", id).Error
	return commodity_info, err
}

func (c *CommodityDbService) Update(commodity_info *model.CommodityInfo) (*model.CommodityInfo, error) {
	commodity_info.UpdatedAt = time.Now().Unix()
	result := database.GormDB.Where("id = ?", commodity_info.ID).Updates(commodity_info)
	return commodity_info, result.Error
}
func (c *CommodityDbService) Delete(user_id, id int64) error {
	commodity_info, err := c.GetById(user_id, id)
	if err != nil {
		return fmt.Errorf("找不到该商品")
	}
	result := database.GormDB.Delete(&commodity_info)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *CommodityDbService) GenerateSnowflakeID() int64 {
	return u.node.Generate().Int64()
}
