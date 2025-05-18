package dto

// PurchaseInboundDTO 进货单数据传输对象
type PurchaseInboundDTO struct {
	CommodityID    int64    `json:"commodityId" binding:"required"`
	Quantity       *int     `json:"quantity" binding:"required"`
	Price          *float64 `json:"price" binding:"required"`
	Specifications *string  `gorm:"type:varchar(255);not null;default:'';comment:商品规格" json:"specifications"`
	OperatorID     int64    `json:"operatorId"`
	Remark         *string  `json:"remark"`
}

// PurchaseReturnDTO 退货单数据传输对象
type PurchaseReturnDTO struct {
	CommodityID    int64    `json:"commodityId" binding:"required"`
	Quantity       *int     `json:"quantity" binding:"required"`
	Price          *float64 `json:"price" binding:"required"`
	Specifications *string  `gorm:"type:varchar(255);not null;default:'';comment:商品规格" json:"specifications"`
	Reason         *string  `json:"reason"`
	OperatorID     int64    `json:"operatorId"`
	Remark         *string  `json:"remark"`
}
