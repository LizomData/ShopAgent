package po

import (
	"time"
)

// PurchaseReturn 退货单
type PurchaseReturn struct {
	ID             int64     `gorm:"primaryKey;column:id;comment:'退货单ID'"`
	CommodityID    int64     `gorm:"column:commodity_id;not null;comment:'商品ID'"`
	Quantity       *int      `gorm:"column:quantity;not null;comment:'退货数量'"`
	Price          *float64  `gorm:"column:price;type:decimal(10,2);not null;comment:'退货单价'"`
	Specifications *string   `gorm:"column:specifications;type:varchar(255);not null;default:'';comment:商品规格"`
	ReturnTime     time.Time `gorm:"column:return_time;type:datetime;default:CURRENT_TIMESTAMP;comment:'退货时间'"`
	OperatorID     int64     `gorm:"column:operator_id;comment:'操作人ID'"`
	SupplierID     int64     `gorm:"column:supplier_id;comment:'供应商ID'"`
	Remark         *string   `gorm:"column:remark;type:varchar(255);comment:'备注'"`
}

// TableName 指定表名
func (PurchaseReturn) TableName() string {
	return "purchase_return"
}
