package vo

import "time"

// PurchaseReturnVO 退货单视图对象
type PurchaseReturnVO struct {
	ID            int64     `json:"id"`
	CommodityID   int64     `json:"commodityId"`
	CommodityName *string   `json:"commodityName"`
	Quantity      *int      `json:"quantity"`
	Price         *float64  `json:"price"`
	TotalAmount   *float64  `json:"totalAmount"`
	Reason        *string   `json:"reason"`
	ReturnTime    time.Time `json:"returnTime"`
	OperatorID    int64     `json:"operatorId"`
	Remark        *string   `json:"remark"`
	SupplierName  *string   `json:"supplierName"`
	SupplierId    int64     `json:"supplierId"`
}
