package po

import "time"

// PurchaseInbound 进货单
type PurchaseInbound struct {
	ID             int64     `gorm:"primaryKey;column:id;comment:'进货单ID'" `
	CommodityID    int64     `gorm:"column:commodity_id;not null;comment:'商品ID'"`
	Quantity       *int      `gorm:"column:quantity;not null;comment:'进货数量'"`
	Price          *float64  `gorm:"column:price;type:decimal(10,2);default:0;not null;comment:'进货单价'"`
	Specifications *string   `gorm:"type:varchar(255);not null;default:'';comment:商品规格" json:"specifications"`
	InboundTime    time.Time `gorm:"column:inbound_time;type:datetime;default:CURRENT_TIMESTAMP;comment:'进货时间'"`
	OperatorID     int64     `gorm:"column:operator_id;comment:'操作人ID'"`
	Remark         *string   `gorm:"column:remark;type:varchar(255);comment:'备注'"`
}

// TableName 指定表名
func (PurchaseInbound) TableName() string {
	return "purchase_inbound"
}
