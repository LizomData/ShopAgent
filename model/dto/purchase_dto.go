package dto

// PurchaseInboundDTO 进货单数据传输对象
type PurchaseInboundDTO struct {
	CommodityID    int64    `json:"commodityId,string" binding:"required"`
	Quantity       *int     `json:"quantity" binding:"required"`
	Price          *float64 `json:"price" binding:"required"`
	Specifications *string  `json:"specifications" binding:"required"`
	Remark         *string  `json:"remark"`
	SupplierID     int64    `json:"supplierId,string" binding:"required"`
}

// PurchaseReturnDTO 退货单数据传输对象
type PurchaseReturnDTO struct {
	CommodityID    int64    `json:"commodityId,string" binding:"required"`
	Quantity       *int     `json:"quantity" binding:"required"`
	Price          *float64 `json:"price" binding:"required"`
	Specifications *string  `json:"specifications" binding:"required"`
	Remark         *string  `json:"remark"`
	SupplierID     int64    `json:"supplierId,string" binding:"required"`
}

// DeleteInboundDTO 删除进货单
type DeleteInboundDTO struct {
	InboundID int64 `json:"inboundID,string" binding:"required"`
}

// DeleteReturnDTO 删除退货单
type DeleteReturnDTO struct {
	ReturnID int64 `json:"returnID,string" binding:"required"`
}

// UpdateInboundDTO 修改进货单
type UpdateInboundDTO struct {
	InboundID      int64    `json:"inboundId,string" binding:"required"`
	CommodityID    int64    `json:"commodityId,string" binding:"required"`
	Quantity       *int     `json:"quantity" binding:"required"`
	Price          *float64 `json:"price" binding:"required"`
	Specifications *string  `json:"specifications" binding:"required"`
	Remark         *string  `json:"remark"`
	SupplierID     int64    `json:"supplierId,string" binding:"required"`
}

// UpdateReturnDTO 修改退货单
type UpdateReturnDTO struct {
	ReturnID       int64    `json:"returnId,string" binding:"required"`
	CommodityID    int64    `json:"commodityId,string" binding:"required"`
	Quantity       *int     `json:"quantity" binding:"required"`
	Price          *float64 `json:"price" binding:"required"`
	Specifications *string  `json:"specifications"`
	Reason         *string  `json:"reason"`
	Remark         *string  `json:"remark"`
	SupplierID     int64    `json:"supplierId,string" binding:"required"`
}
