package dto

// PurchaseInboundDTO 进货单数据传输对象
type PurchaseInboundDTO struct {
	CommodityID    int64    `json:"commodityId,string" binding:"required"`
	Quantity       *int     `json:"quantity" binding:"required"`
	Price          *float64 `json:"price" binding:"required"`
	Specifications *string  `gorm:"type:varchar(255);not null;default:'';comment:商品规格" json:"specifications" binding:"required"`
	Remark         *string  `json:"remark"`
	SupplierID     int64    `json:"supplierId,string" binding:"required"`
}

// PurchaseReturnDTO 退货单数据传输对象
type PurchaseReturnDTO struct {
	CommodityID    int64    `json:"commodityId,string" binding:"required"`
	Quantity       *int     `json:"quantity" binding:"required"`
	Price          *float64 `json:"price" binding:"required"`
	Specifications *string  `gorm:"type:varchar(255);not null;default:'';comment:商品规格" json:"specifications"`
	Reason         *string  `json:"reason"`
	Remark         *string  `json:"remark"`
	SupplierID     int64    `json:"supplierId,string" binding:"required"`
}
