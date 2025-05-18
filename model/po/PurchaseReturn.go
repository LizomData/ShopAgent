package po

import (
	"time"
)

// PurchaseReturn 退货单
type PurchaseReturn struct {
	ID          int64     `gorm:"primaryKey;column:id;comment:'退货单ID'"`
	CommodityID int64     `gorm:"column:commodity_id;not null;comment:'商品ID'"`
	Quantity    *int      `gorm:"column:quantity;not null;comment:'退货数量'"`
	Price       *float64  `gorm:"column:price;type:decimal(10,2);not null;comment:'退货单价'"`
	Reason      *string   `gorm:"column:reason;type:varchar(255);comment:'退货原因'"`
	ReturnTime  time.Time `gorm:"column:return_time;type:datetime;default:CURRENT_TIMESTAMP;comment:'退货时间'"`
	OperatorID  int64     `gorm:"column:operator_id;comment:'操作人ID'"`
	Remark      *string   `gorm:"column:remark;type:varchar(255);comment:'备注'"`
}

// TableName 指定表名
func (PurchaseReturn) TableName() string {
	return "purchase_return"
}
