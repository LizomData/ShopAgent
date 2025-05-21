package commodityDbService

import (
	"ShopAgent/model"
	"ShopAgent/util/database"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"
)

// Instance 全局单例实例
var Instance = GetInstance()

// GetInstance 获取CommodityDbService的单例实例
// 返回: *CommodityDbService 服务实例
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

// CommodityDbService 商品数据库服务
type CommodityDbService struct {
	node *snowflake.Node // 雪花算法节点
}

// Create 创建新商品
// commodity_info: 商品信息
// 返回: 创建后的商品信息和可能的错误
func (u *CommodityDbService) Create(commodity_info *model.CommodityInfo) (*model.CommodityInfo, error) {
	commodity_info.ID = u.GenerateSnowflakeID()
	commodity_info.CreatedAt = time.Now().Unix()
	commodity_info.UpdatedAt = time.Now().Unix()
	result := database.GormDB.Create(&commodity_info)
	return commodity_info, result.Error
}

// Query 查询商品列表
// user_id: 用户ID
// page: 页码
// size: 每页数量
// 返回: 商品列表和可能的错误
func (s *CommodityDbService) Query(user_id int64, page, size int, name *string) ([]model.CommodityInfo, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}

	offset := (page - 1) * size
	var commodity_infos []model.CommodityInfo
	var total int64 // 新增总记录数变量

	// 先查询总记录数
	if err := database.GormDB.Model(&model.CommodityInfo{}).
		Where("user_id = ?", user_id).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query := database.GormDB.
		Where("user_id = ?", user_id) // 固定条件

	// 仅在 name 非空时添加条件
	if *name != "" {
		query = query.Where("name = ?", name)
	}
	err := query.
		Offset(offset).
		Limit(size).
		Find(&commodity_infos).Error

	return commodity_infos, total, err
}

// GetById 根据ID获取商品信息
// user_id: 用户ID
// id: 商品ID
// 返回: 商品信息和可能的错误
func (s *CommodityDbService) GetById(user_id, id int64) (model.CommodityInfo, error) {
	var commodity_info model.CommodityInfo
	err := database.GormDB.
		Where("user_id = ?", user_id).
		First(&commodity_info, "id = ?", id).Error
	return commodity_info, err
}

// Update 更新商品信息
// commodity_info: 需要更新的商品信息
// 返回: 更新后的商品信息和可能的错误
func (c *CommodityDbService) Update(user_id int64, commodity_info *model.CommodityInfo) (*model.CommodityInfo, error) {

	_commodity_info, err := c.GetById(user_id, commodity_info.ID)
	if err != nil {
		return nil, fmt.Errorf("查询不到该商品")
	}
	commodity_info.UserID = _commodity_info.UserID
	commodity_info.CreatedAt = _commodity_info.CreatedAt

	commodity_info.UpdatedAt = time.Now().Unix()
	result := database.GormDB.Where("id = ?", commodity_info.ID).Updates(commodity_info)
	return commodity_info, result.Error
}

// Delete 删除商品
// user_id: 用户ID
// id: 商品ID
// 返回: 删除操作的错误信息
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

// GenerateSnowflakeID 生成雪花算法ID
// 返回: 生成的唯一ID
func (u *CommodityDbService) GenerateSnowflakeID() int64 {
	return u.node.Generate().Int64()
}

// UpdateStock 更新商品库存
// commodityID: 商品ID
// quantity: 库存变化量（正数增加，负数减少）
// 返回: 更新操作的错误信息
func (c *CommodityDbService) UpdateStock(commodityID int64, quantity int) error {
	// 更新商品库存
	result := database.GormDB.Model(&model.CommodityInfo{}).Where("id = ?", commodityID).UpdateColumn("quantity + ?", quantity)
	return result.Error
}
