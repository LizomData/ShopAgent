package purchaseDbService

import (
	"ShopAgent/model"
	"ShopAgent/model/dto"
	"ShopAgent/model/po"
	"ShopAgent/model/vo"
	"ShopAgent/util/database"
	"ShopAgent/util/database/commodityDbService"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"time"
)

// Instance 全局单例实例
var Instance = GetInstance()

// GetInstance 获取PurchaseDbService的单例实例
// 返回: *PurchaseDbService 服务实例
func GetInstance() *purchase_db_service {
	// 初始化节点（确保每个服务实例的 nodeID 唯一）
	nodeID := int64(6)
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		fmt.Println("初始化雪花节点失败: " + err.Error())
		return nil
	}

	// 初始化 CommodityRepo
	commodityRepo := commodityDbService.GetInstance()

	return &purchase_db_service{
		node,
		commodityRepo,
	}
}

// PurchaseDbService 进货和退货服务
type purchase_db_service struct {
	node          *snowflake.Node                        // 雪花算法节点
	CommodityRepo *commodityDbService.CommodityDbService // 商品仓库
}

// CreateInbound 创建进货单
func (s *purchase_db_service) CreateInbound(dto *dto.PurchaseInboundDTO, user_id int64) (vo.PurchaseInboundVO, error) {

	// 创建一个空的 VO 对象
	var inboundVO vo.PurchaseInboundVO
	fmt.Printf("dto: %v\n", dto)

	// 1、先根据商品ID获取商品数据
	commodity, _ := s.CommodityRepo.GetById(user_id, dto.CommodityID)
	inbound := &po.PurchaseInbound{
		CommodityID:    dto.CommodityID,
		Quantity:       dto.Quantity,
		Price:          commodity.Price,
		Specifications: commodity.Specifications,
		InboundTime:    time.Now(),
		OperatorID:     user_id,
		Remark:         dto.Remark,
	}

	fmt.Printf("获取到的 commodity 信息:\n%+v\n", commodity)

	if err := database.GormDB.Create(inbound).Error; err != nil {
		return inboundVO, err
	}

	if err := s.UpdateStock(dto.CommodityID, dto.Quantity); err != nil {
		return inboundVO, err
	}

	inboundVO = vo.PurchaseInboundVO{
		ID:            inbound.ID,
		CommodityID:   inbound.CommodityID,
		CommodityName: commodity.Name,
		Quantity:      inbound.Quantity,
		Price:         inbound.Price,
		TotalAmount:   new(float64),
		InboundTime:   inbound.InboundTime,
		OperatorId:    user_id,
		Remark:        inbound.Remark,
	}

	return inboundVO, nil
}

// CreateReturn 创建退货单
func (s *purchase_db_service) CreateReturn(dto *dto.PurchaseReturnDTO, user_id int64) (vo.PurchaseReturnVO, error) {

	var returnVO vo.PurchaseReturnVO
	commodity, _ := s.CommodityRepo.GetById(user_id, dto.CommodityID)
	fmt.Printf("获取到的 commodity 信息:\n%+v\n", commodity)

	returnOrder := &po.PurchaseReturn{
		CommodityID: dto.CommodityID,
		Quantity:    dto.Quantity,
		Price:       dto.Price,
		Reason:      dto.Reason,
		ReturnTime:  time.Now(),
		OperatorID:  user_id,
		Remark:      dto.Remark,
	}

	if err := database.GormDB.Create(returnOrder).Error; err != nil {
		return returnVO, err
	}

	if err := s.UpdateStock(dto.CommodityID, dto.Quantity); err != nil {
		return returnVO, err
	}

	returnVO = vo.PurchaseReturnVO{
		ID:            returnOrder.ID,
		CommodityID:   returnOrder.CommodityID,
		CommodityName: commodity.Name,
		Quantity:      returnOrder.Quantity,
		Price:         returnOrder.Price,
		TotalAmount:   new(float64),
		ReturnTime:    returnOrder.ReturnTime,
		OperatorID:    user_id,
		Remark:        returnOrder.Remark,
	}

	return returnVO, nil
}

// GetInboundList 获取进货单列表
func (s *purchase_db_service) GetInboundList(page, pageSize int, user_id int64) ([]vo.PurchaseInboundVO, int64, error) {
	var inbounds []po.PurchaseInbound
	var total int64

	err := database.GormDB.Model(&po.PurchaseInbound{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = database.GormDB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&inbounds).Error
	if err != nil {
		return nil, 0, err
	}

	vos := make([]vo.PurchaseInboundVO, len(inbounds))
	for i, inbound := range inbounds {

		commodity, _ := s.CommodityRepo.GetById(user_id, inbound.CommodityID)
		fmt.Printf("获取到的 commodity 信息:\n%+v\n", commodity.Name)

		vos[i] = vo.PurchaseInboundVO{
			ID:            inbound.ID,
			CommodityID:   inbound.CommodityID,
			CommodityName: commodity.Name,
			Quantity:      inbound.Quantity,
			Price:         inbound.Price,
			TotalAmount:   new(float64),
			InboundTime:   inbound.InboundTime,
			OperatorId:    user_id,
			Remark:        inbound.Remark,
		}
		*vos[i].TotalAmount = float64(*inbound.Quantity) * *inbound.Price
	}

	return vos, total, nil
}

// GetReturnList 获取退货单列表
func (s *purchase_db_service) GetReturnList(page, pageSize int, user_id int64) ([]vo.PurchaseReturnVO, int64, error) {
	var returns []po.PurchaseReturn
	var total int64

	err := database.GormDB.Model(&po.PurchaseReturn{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = database.GormDB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&returns).Error
	if err != nil {
		return nil, 0, err
	}

	vos := make([]vo.PurchaseReturnVO, len(returns))
	for i, ret := range returns {
		var commodity model.CommodityInfo
		//// 查询商品信息
		//err := database.GormDB.First(&commodity, "id = ?", ret.CommodityID).Error
		//if err != nil {
		//	return nil, 0, err
		//}

		commodity, _ = s.CommodityRepo.GetById(user_id, ret.CommodityID)

		vos[i] = vo.PurchaseReturnVO{
			ID:            ret.ID,
			CommodityID:   ret.CommodityID,
			CommodityName: commodity.Name,
			Quantity:      ret.Quantity,
			Price:         ret.Price,
			TotalAmount:   new(float64),
			Reason:        ret.Reason,
			ReturnTime:    ret.ReturnTime,
			OperatorID:    user_id,
			Remark:        ret.Remark,
		}
		*vos[i].TotalAmount = float64(*ret.Quantity) * *ret.Price
	}

	return vos, total, nil
}

// UpdateStock 更新商品库存
func (s *purchase_db_service) UpdateStock(commodityID int64, quantity *int) error {
	result := database.GormDB.Model(&model.CommodityInfo{}).Where("id = ?", commodityID).UpdateColumn("quantity", gorm.Expr("quantity + ?", quantity))
	return result.Error
}
