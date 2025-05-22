package vo

import "time"

// PurchaseInboundVO 进货单视图对象
type PurchaseInboundVO struct {
	ID             int64     `json:"id"`
	CommodityID    int64     `json:"commodityId"`
	CommodityName  *string   `json:"commodityName"`
	Quantity       *int      `json:"quantity"`
	Price          *float64  `json:"price"`
	TotalAmount    *float64  `json:"totalAmount"`
	InboundTime    time.Time `json:"inboundTime"`
	OperatorId     int64     `json:"operatorName"`
	Remark         *string   `json:"remark"`
	SupplierName   *string   `json:"supplierName"`
	SupplierId     int64     `json:"supplierId"`
	Specifications *string   `json:"specifications"`
}
