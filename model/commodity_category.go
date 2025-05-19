package model

type CommodityCategory struct {
	ID          int64   `gorm:"type:bigint;primary_key;comment:商品类别id" json:"id,string"`
	UserID      int64   `gorm:"type:bigint;not null;comment:关联用户ID" json:"-"`
	Name        *string `gorm:"type:varchar(255);unique;not null;default:'';comment:类别名称" json:"name" binding:"required"`
	Description *string `gorm:"type:varchar(255);not null;default:'';comment:类别描述" json:"description" binding:"required"`
	CardImage   *string `gorm:"type:varchar(255);not null;default:'';comment:类别图片" json:"card_image" binding:"required"`
	CreatedAt   int64   `gorm:"type:bigint;comment:创建时间" json:"created_at"`
	UpdatedAt   int64   `gorm:"type:bigint;comment:更新时间" json:"updated_at"`
}
