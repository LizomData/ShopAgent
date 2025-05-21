package imageUploadDbService

import (
	"ShopAgent/model"
	"ShopAgent/util/database"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"
)

// Instance 全局单例实例
var Instance = GetInstance()

// GetInstance 获取ImageUploadDbService的单例实例
// 返回: *ImageUploadDbService 服务实例
func GetInstance() *ImageUploadDbService {
	// 初始化节点（确保每个服务实例的 nodeID 唯一）
	nodeID := int64(7)
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		fmt.Println("初始化雪花节点失败: " + err.Error())
		return nil
	}
	return &ImageUploadDbService{node}
}

// ImageUploadDbService 商品数据库服务
type ImageUploadDbService struct {
	node *snowflake.Node // 雪花算法节点
}

func (u *ImageUploadDbService) Create(user_id int64, image_upload_record *model.ImageUploadRecord) (*model.ImageUploadRecord, error) {
	image_upload_record.ID = u.GenerateSnowflakeID()
	image_upload_record.UserID = user_id
	image_upload_record.CreatedAt = time.Now().Unix()
	image_upload_record.UpdatedAt = time.Now().Unix()
	result := database.GormDB.Create(&image_upload_record)
	return image_upload_record, result.Error
}

// Query 查询商品列表
// user_id: 用户ID
// page: 页码
// size: 每页数量
// 返回: 商品列表和可能的错误
func (s *ImageUploadDbService) Query(user_id int64, page, size int) ([]model.ImageUploadRecord, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}

	offset := (page - 1) * size
	var image_upload_records []model.ImageUploadRecord
	var total int64 // 新增总记录数变量

	// 先查询总记录数
	if err := database.GormDB.Model(&model.ImageUploadRecord{}).
		Where("user_id = ?", user_id).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := database.GormDB.
		Where("user_id = ?", user_id).
		Offset(offset).
		Limit(size).
		Find(&image_upload_records).Error

	return image_upload_records, total, err
}

// GetById 根据ID获取商品信息
// user_id: 用户ID
// id: 商品ID
// 返回: 商品信息和可能的错误
func (s *ImageUploadDbService) GetById(user_id, id int64) (model.ImageUploadRecord, error) {
	var image_upload_record model.ImageUploadRecord
	err := database.GormDB.
		Where("user_id = ?", user_id).
		First(&image_upload_record, "id = ?", id).Error
	return image_upload_record, err
}

// Update 更新商品信息
// image_upload_record: 需要更新的商品信息
// 返回: 更新后的商品信息和可能的错误
func (c *ImageUploadDbService) Update(user_id int64, image_upload_record *model.ImageUploadRecord) (*model.ImageUploadRecord, error) {

	_image_upload_record, err := c.GetById(user_id, image_upload_record.ID)
	if err != nil {
		return nil, fmt.Errorf("查询不到该商品")
	}
	image_upload_record.UserID = _image_upload_record.UserID
	image_upload_record.CreatedAt = _image_upload_record.CreatedAt

	image_upload_record.UpdatedAt = time.Now().Unix()
	result := database.GormDB.Where("id = ?", image_upload_record.ID).Updates(image_upload_record)
	return image_upload_record, result.Error
}

// Delete 删除商品
// user_id: 用户ID
// id: 商品ID
// 返回: 删除操作的错误信息
func (c *ImageUploadDbService) Delete(user_id, id int64) error {
	image_upload_record, err := c.GetById(user_id, id)
	if err != nil {
		return fmt.Errorf("找不到该商品")
	}
	result := database.GormDB.Delete(&image_upload_record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GenerateSnowflakeID 生成雪花算法ID
// 返回: 生成的唯一ID
func (u *ImageUploadDbService) GenerateSnowflakeID() int64 {
	return u.node.Generate().Int64()
}

// UpdateStock 更新商品库存
// commodityID: 商品ID
// quantity: 库存变化量（正数增加，负数减少）
// 返回: 更新操作的错误信息
func (c *ImageUploadDbService) UpdateStock(commodityID int64, quantity int) error {
	// 更新商品库存
	result := database.GormDB.Model(&model.ImageUploadRecord{}).Where("id = ?", commodityID).UpdateColumn("quantity + ?", quantity)
	return result.Error
}
