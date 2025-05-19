package model

type Supplier struct {
	ID             int64   `gorm:"type:bigint;primary_key" json:"id,string"`
	UserID         int64   `gorm:"type:bigint;not null;comment:关联用户ID" json:"-"`
	Name           *string `gorm:"type:varchar(255);not null;default:''" json:"name" binding:"required"`
	PersonInCharge *string `gorm:"type:varchar(255);not null;default:'';comment:负责人" json:"person_in_charge" binding:"required"`
	ContactNumber  *string `gorm:"type:varchar(255);not null;default:'';comment:负责人电话" json:"contact_number" binding:"required"`
	Address        *string `gorm:"type:varchar(255);not null;default:'';comment:地址" json:"address" binding:"required"`
	CreatedAt      int64   `gorm:"type:bigint;comment:创建时间" json:"created_at"`
	UpdatedAt      int64   `gorm:"type:bigint;comment:更新时间" json:"updated_at"`
}
