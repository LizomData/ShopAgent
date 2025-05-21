package purchaseDbService

import (
	"ShopAgent/model"
	"ShopAgent/model/dto"
	"ShopAgent/model/po"
	"ShopAgent/model/vo"
	"ShopAgent/util/database"
	"ShopAgent/util/database/commodityDbService"
	"ShopAgent/util/database/supplierDbService"
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"time"
)

// Instance 全局单例实例
var Instance = GetInstance()

// GetInstance 获取PurchaseDbService的单例实例
// 返回: *PurchaseDbService 服务实例
func GetInstance() *PurchaseDbService {
	// 初始化节点（确保每个服务实例的 nodeID 唯一）
	nodeID := int64(6)
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		fmt.Println("初始化雪花节点失败: " + err.Error())
		return nil
	}

	// 初始化 CommodityRepo
	commodityRepo := commodityDbService.GetInstance()

	// 初始化 SupplierRepo
	supplierRepo := supplierDbService.GetInstance()

	return &PurchaseDbService{
		node,
		commodityRepo,
		supplierRepo,
	}
}

// PurchaseDbService 进货和退货服务
type PurchaseDbService struct {
	node          *snowflake.Node                        // 雪花算法节点
	CommodityRepo *commodityDbService.CommodityDbService // 商品仓库
	SupplierRepo  *supplierDbService.SupplierDbService   // 供应商仓库
}

// CreateInbound 创建进货单
func (s *PurchaseDbService) CreateInbound(dto *dto.PurchaseInboundDTO, userId int64) (vo.PurchaseInboundVO, error) {

	// 创建一个空的 VO 对象
	var inboundVO vo.PurchaseInboundVO

	// 1. 查询供货商数据
	supplier, _ := s.SupplierRepo.GetById(userId, dto.SupplierID)

	// 检查供应商ID是否为空
	if supplier.ID == 0 {
		return inboundVO, fmt.Errorf("供应商不存在")
	}

	// 2. 先根据商品ID获取商品数据
	commodity, _ := s.CommodityRepo.GetById(userId, dto.CommodityID)

	if commodity.ID == 0 {
		return inboundVO, fmt.Errorf("商品不存在")
	}
	inbound := &po.PurchaseInbound{
		ID:             s.node.Generate().Int64(),
		CommodityID:    dto.CommodityID,
		Quantity:       dto.Quantity,
		Price:          commodity.Price,
		Specifications: commodity.Specifications,
		InboundTime:    time.Now(),
		OperatorID:     userId,
		SupplierID:     dto.SupplierID,
		Remark:         dto.Remark,
	}

	// 3. 插入数据
	if err := database.GormDB.Create(inbound).Error; err != nil {
		return inboundVO, err
	}

	// 4. 更新商品数量 已经改成触发器
	//if err := s.UpdateStock(dto.CommodityID, dto.Quantity); err != nil {
	//	return inboundVO, err
	//}

	// 5. 构建视图返回对象
	inboundVO = vo.PurchaseInboundVO{
		ID:            inbound.ID,
		CommodityID:   inbound.CommodityID,
		CommodityName: commodity.Name,
		Quantity:      inbound.Quantity,
		Price:         inbound.Price,
		TotalAmount:   new(float64),
		InboundTime:   inbound.InboundTime,
		OperatorId:    userId,
		SupplierName:  supplier.Name,
		Remark:        inbound.Remark,
	}

	*inboundVO.TotalAmount = float64(*inboundVO.Quantity) * *inboundVO.Price

	return inboundVO, nil
}

// CreateReturn 创建退货单
func (s *PurchaseDbService) CreateReturn(dto *dto.PurchaseReturnDTO, user_id int64) (vo.PurchaseReturnVO, error) {

	var returnVO vo.PurchaseReturnVO

	// 查询商品
	commodity, _ := s.CommodityRepo.GetById(user_id, dto.CommodityID)

	// 查询供应商
	supplier, _ := s.SupplierRepo.GetById(user_id, dto.SupplierID)

	// 检查供应商ID是否为空
	if supplier.ID == 0 {
		return returnVO, fmt.Errorf("供应商不存在")
	}

	// 检查库存是否足够
	if commodity.Quantity == nil || *commodity.Quantity < *dto.Quantity {
		return returnVO, fmt.Errorf("商品库存不足，当前库存: %d, 退货数量: %d", *commodity.Quantity, *dto.Quantity)
	}

	returnOrder := &po.PurchaseReturn{
		ID:          s.node.Generate().Int64(),
		CommodityID: dto.CommodityID,
		Quantity:    dto.Quantity,
		Price:       dto.Price,
		Reason:      dto.Reason,
		ReturnTime:  time.Now(),
		OperatorID:  user_id,
		SupplierID:  dto.SupplierID,
		Remark:      dto.Remark,
	}

	if err := database.GormDB.Create(returnOrder).Error; err != nil {
		return returnVO, err
	}

	// 在这里进行商品数量的操作，已经改成触发器
	//if err := s.UpdateStock(dto.CommodityID, dto.Quantity); err != nil {
	//	return returnVO, err
	//}

	returnVO = vo.PurchaseReturnVO{
		ID:            returnOrder.ID,
		CommodityID:   returnOrder.CommodityID,
		CommodityName: commodity.Name,
		Quantity:      returnOrder.Quantity,
		Price:         returnOrder.Price,
		TotalAmount:   new(float64),
		ReturnTime:    returnOrder.ReturnTime,
		OperatorID:    user_id,
		SupplierName:  supplier.Name,
		Remark:        returnOrder.Remark,
	}

	*returnVO.TotalAmount = float64(*returnVO.Quantity) * *returnVO.Price

	return returnVO, nil
}

// GetInboundList 获取进货单列表
func (s *PurchaseDbService) GetInboundList(page, pageSize int, user_id int64) ([]vo.PurchaseInboundVO, int64, error) {
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

		// 查询供应商
		supplier, _ := s.SupplierRepo.GetById(user_id, inbound.SupplierID)

		vos[i] = vo.PurchaseInboundVO{
			ID:            inbound.ID,
			CommodityID:   inbound.CommodityID,
			CommodityName: commodity.Name,
			Quantity:      inbound.Quantity,
			Price:         inbound.Price,
			TotalAmount:   new(float64),
			InboundTime:   inbound.InboundTime,
			OperatorId:    user_id,
			SupplierName:  supplier.Name,
			Remark:        inbound.Remark,
		}
		*vos[i].TotalAmount = float64(*inbound.Quantity) * *inbound.Price
	}

	return vos, total, nil
}

// GetReturnList 获取退货单列表
func (s *PurchaseDbService) GetReturnList(page, pageSize int, user_id int64) ([]vo.PurchaseReturnVO, int64, error) {
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

		commodity, _ = s.CommodityRepo.GetById(user_id, ret.CommodityID)
		supplier, _ := s.SupplierRepo.GetById(user_id, ret.SupplierID)

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
			SupplierName:  supplier.Name,
			Remark:        ret.Remark,
		}
		*vos[i].TotalAmount = float64(*ret.Quantity) * *ret.Price
	}

	return vos, total, nil
}

// UpdateStock 更新商品库存
func (s *PurchaseDbService) UpdateStock(commodityID int64, quantity *int) error {
	result := database.GormDB.Model(&model.CommodityInfo{}).Where("id = ?", commodityID).UpdateColumn("quantity", gorm.Expr("quantity + ?", quantity))
	return result.Error
}

// DeleteInbound 删除进货单
func (s *PurchaseDbService) DeleteInbound(userId int64, inboundId int64) error {
	var inbound po.PurchaseInbound

	// 查询进货单是否存在
	if err := database.GormDB.Where("id = ? AND operator_id = ?", inboundId, userId).First(&inbound).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("进货单不存在")
		}
		return err
	}

	// 开启事务
	return database.GormDB.Transaction(func(tx *gorm.DB) error {
		// 删除进货单
		if err := tx.Delete(&inbound).Error; err != nil {
			return err
		}

		// 更新商品库存（减去进货数量）
		negativeQuantity := -(*inbound.Quantity)
		if err := s.UpdateStock(inbound.CommodityID, &negativeQuantity); err != nil {
			return err
		}

		return nil
	})
}

// DeleteReturn 删除退货单
func (s *PurchaseDbService) DeleteReturn(userId int64, returnId int64) error {
	var returnOrder po.PurchaseReturn

	// 查询退货单是否存在
	if err := database.GormDB.Where("id = ? AND operator_id = ?", returnId, userId).First(&returnOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("退货单不存在")
		}
		return err
	}

	// 开启事务
	return database.GormDB.Transaction(func(tx *gorm.DB) error {
		// 删除退货单
		if err := tx.Delete(&returnOrder).Error; err != nil {
			return err
		}

		// 更新商品库存（加回退货数量）
		if err := s.UpdateStock(returnOrder.CommodityID, returnOrder.Quantity); err != nil {
			return err
		}

		return nil
	})
}

// UpdateInbound 修改进货单
func (s *PurchaseDbService) UpdateInbound(userId int64, dto *dto.UpdateInboundDTO) error {
	var inbound po.PurchaseInbound

	// 查询进货单是否存在
	if err := database.GormDB.Where("id = ? AND operator_id = ?", dto.InboundID, userId).First(&inbound).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("进货单不存在")
		}
		return err
	}

	// 查询供应商
	supplier, _ := s.SupplierRepo.GetById(userId, dto.SupplierID)
	if supplier.ID == 0 {
		return fmt.Errorf("供应商不存在")
	}

	// 查询商品
	commodity, _ := s.CommodityRepo.GetById(userId, dto.CommodityID)
	if commodity.ID == 0 {
		return fmt.Errorf("商品不存在")
	}

	// 开启事务
	return database.GormDB.Transaction(func(tx *gorm.DB) error {
		// 计算库存差值
		quantityDiff := *dto.Quantity - *inbound.Quantity

		// 更新进货单信息
		inbound.CommodityID = dto.CommodityID
		inbound.Quantity = dto.Quantity
		inbound.Price = dto.Price
		inbound.Specifications = dto.Specifications
		inbound.SupplierID = dto.SupplierID
		inbound.Remark = dto.Remark

		// 更新进货单
		if err := tx.Save(&inbound).Error; err != nil {
			return err
		}

		// 更新商品库存
		if err := s.UpdateStock(inbound.CommodityID, &quantityDiff); err != nil {
			return err
		}

		return nil
	})
}

// UpdateReturn 修改退货单
func (s *PurchaseDbService) UpdateReturn(userId int64, dto *dto.UpdateReturnDTO) error {
	var returnOrder po.PurchaseReturn

	// 查询退货单是否存在
	if err := database.GormDB.Where("id = ? AND operator_id = ?", dto.ReturnID, userId).First(&returnOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("退货单不存在")
		}
		return err
	}

	// 查询供应商
	supplier, _ := s.SupplierRepo.GetById(userId, dto.SupplierID)
	if supplier.ID == 0 {
		return fmt.Errorf("供应商不存在")
	}

	// 查询商品
	commodity, _ := s.CommodityRepo.GetById(userId, dto.CommodityID)
	if commodity.ID == 0 {
		return fmt.Errorf("商品不存在")
	}

	// 计算库存差值
	quantityDiff := *dto.Quantity - *returnOrder.Quantity

	// 检查库存是否足够
	if commodity.Quantity == nil || *commodity.Quantity < quantityDiff {
		return fmt.Errorf("商品库存不足，当前库存: %d, 需要的库存变化: %d", *commodity.Quantity, quantityDiff)
	}

	// 开启事务
	return database.GormDB.Transaction(func(tx *gorm.DB) error {
		// 更新退货单信息
		returnOrder.CommodityID = dto.CommodityID
		returnOrder.Quantity = dto.Quantity
		returnOrder.Price = dto.Price
		returnOrder.Reason = dto.Reason
		returnOrder.SupplierID = dto.SupplierID
		returnOrder.Remark = dto.Remark

		// 更新退货单
		if err := tx.Save(&returnOrder).Error; err != nil {
			return err
		}

		// 更新商品库存（退货会减少库存）
		negativeQuantityDiff := -quantityDiff
		if err := s.UpdateStock(returnOrder.CommodityID, &negativeQuantityDiff); err != nil {
			return err
		}

		return nil
	})
}
