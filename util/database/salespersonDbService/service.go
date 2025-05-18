package salespersonDbService

import (
	"ShopAgent/model"
	"ShopAgent/util/database"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"
)

var Instance = GetInstance()

func GetInstance() *SalespersonDbService {
	// 初始化节点（确保每个服务实例的 nodeID 唯一）
	nodeID := int64(5)
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		fmt.Println("初始化雪花节点失败: " + err.Error())
		return nil
	}
	return &SalespersonDbService{node}
}

type SalespersonDbService struct {
	node *snowflake.Node
}

func (u *SalespersonDbService) Create(salesperson *model.SalesPerson) (*model.SalesPerson, error) {
	salesperson.ID = u.GenerateSnowflakeID()
	salesperson.CreatedAt = time.Now().Unix()
	salesperson.UpdatedAt = time.Now().Unix()
	result := database.GormDB.Create(&salesperson)
	return salesperson, result.Error
}
func (s *SalespersonDbService) Query(user_id int64, page, size int) ([]model.SalesPerson, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}

	offset := (page - 1) * size
	var salespersons []model.SalesPerson

	err := database.GormDB.
		Where("user_id = ?", user_id).
		Offset(offset).
		Limit(size).
		Find(&salespersons).Error

	return salespersons, err
}

func (s *SalespersonDbService) GetById(user_id, id int64) (model.SalesPerson, error) {
	var salesperson model.SalesPerson
	err := database.GormDB.
		Where("user_id = ?", user_id).
		First(&salesperson, "id = ?", id).Error
	return salesperson, err
}

func (c *SalespersonDbService) Update(user_id int64, salesperson *model.SalesPerson) (*model.SalesPerson, error) {

	_salesperson, err := c.GetById(user_id, salesperson.ID)
	if err != nil {
		return nil, fmt.Errorf("查询不到该销售员")
	}
	salesperson.UserID = _salesperson.UserID
	salesperson.CreatedAt = _salesperson.CreatedAt
	salesperson.UpdatedAt = time.Now().Unix()
	result := database.GormDB.Where("id = ?", salesperson.ID).Updates(salesperson)
	return salesperson, result.Error
}

func (c *SalespersonDbService) Delete(user_id, id int64) error {
	salesperson, err := c.GetById(user_id, id)
	if err != nil {
		return fmt.Errorf("找不到该销售员")
	}
	result := database.GormDB.Delete(&salesperson)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *SalespersonDbService) GenerateSnowflakeID() int64 {
	return u.node.Generate().Int64()
}
