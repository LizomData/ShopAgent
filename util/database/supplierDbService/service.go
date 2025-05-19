package supplierDbService

import (
	"ShopAgent/model"
	"ShopAgent/util/database"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"
)

var Instance = GetInstance()

func GetInstance() *SupplierDbService {
	// 初始化节点（确保每个服务实例的 nodeID 唯一）
	nodeID := int64(3)
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		fmt.Println("初始化雪花节点失败: " + err.Error())
		return nil
	}
	return &SupplierDbService{node}
}

type SupplierDbService struct {
	node *snowflake.Node
}

func (u *SupplierDbService) Create(supplier *model.Supplier) (*model.Supplier, error) {
	supplier.ID = u.GenerateSnowflakeID()
	supplier.CreatedAt = time.Now().Unix()
	supplier.UpdatedAt = time.Now().Unix()
	result := database.GormDB.Create(&supplier)
	return supplier, result.Error
}
func (s *SupplierDbService) Query(user_id int64, page, size int) ([]model.Supplier, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}

	offset := (page - 1) * size
	var suppliers []model.Supplier
	var total int64 // 新增总记录数变量

	// 先查询总记录数
	if err := database.GormDB.Model(&model.Supplier{}).
		Where("user_id = ?", user_id).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 再查询分页数据
	err := database.GormDB.
		Where("user_id = ?", user_id).
		Offset(offset).
		Limit(size).
		Find(&suppliers).Error

	return suppliers, total, err
}

func (s *SupplierDbService) GetById(user_id, id int64) (model.Supplier, error) {
	var supplier model.Supplier
	err := database.GormDB.
		Where("user_id = ?", user_id).
		First(&supplier, "id = ?", id).Error
	return supplier, err
}

func (c *SupplierDbService) Update(user_id int64, supplier *model.Supplier) (*model.Supplier, error) {

	_supplier, err := c.GetById(user_id, supplier.ID)
	if err != nil {
		return nil, fmt.Errorf("查询不到该供应商")
	}
	supplier.UserID = _supplier.UserID
	supplier.CreatedAt = _supplier.CreatedAt
	supplier.UpdatedAt = time.Now().Unix()
	result := database.GormDB.Where("id = ?", supplier.ID).Updates(supplier)
	return supplier, result.Error
}

func (c *SupplierDbService) Delete(user_id, id int64) error {
	supplier, err := c.GetById(user_id, id)
	if err != nil {
		return fmt.Errorf("找不到该供应商")
	}
	result := database.GormDB.Delete(&supplier)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *SupplierDbService) GenerateSnowflakeID() int64 {
	return u.node.Generate().Int64()
}
