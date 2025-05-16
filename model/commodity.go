package model

type CommodityInfo struct {
	ID             int64    `gorm:"primary_key;type:bigint;comment:商品id" json:"id"`
	UserID         int64    `gorm:"not null;comment:关联用户ID" json:"-"`
	Name           *string  `gorm:"type:varchar(255);unique;not null;default:'';comment:商品名称" json:"name" binding:"required"`
	Type           *string  `gorm:"type:varchar(255);not null;default:'';comment:商品类型" json:"type" binding:"required"`
	Specifications *string  `gorm:"type:varchar(255);not null;default:'';comment:商品规格" json:"specifications" binding:"required"`
	Quantity       *int     `gorm:"type:bigint;not null;default:0;comment:商品数量" json:"quantity" binding:"required"`
	Price          *float64 `gorm:"type:double(12,2) unsigned;not null;default:0;comment:商品价格" json:"price" binding:"required"`
	CartImage      *string  `gorm:"type:varchar(255);not null;default:'';comment:商品图片" json:"cart_image" binding:"required"`
	CreatedAt      int64    `gorm:"type:bigint;comment:创建时间" json:"created_at"`
	UpdatedAt      int64    `gorm:"type:bigint;comment:更新时间" json:"updated_at"`
}
