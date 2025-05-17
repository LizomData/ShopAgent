package model

type SalesPerson struct {
	ID             int64    `gorm:"type:bigint;primary_key;comment:销售员id" json:"id"`
	UserID         int64    `gorm:"type:bigint;not null;comment:关联用户ID" json:"-"`
	Name           *string  `gorm:"type:varchar(255);not null;default:'';comment:销售员姓名" json:"name" binding:"required"`
	Gender         *string  `gorm:"type:ENUM('M','F','O','U') NOT NULL DEFAULT 'U';comment:销售员性别（M:男 F:女 O:其他 U:未知）" json:"gender" binding:"required,oneof=M F O U"`
	ContactNumber  *string  `gorm:"type:varchar(255);not null;default:'';comment:销售员电话" json:"contact_number" binding:"required"`
	Address        *string  `gorm:"type:varchar(255);not null;default:'';comment:地址" json:"address" binding:"required"`
	ProfilePicture *string  `gorm:"type:varchar(255);not null;default:'';comment:头像" json:"profile_picture" binding:"required"`
	Salary         *float64 `gorm:"type:double(12,2) unsigned;not null;default:0;comment:薪资" json:"salary" binding:"required"`
	CreatedAt      int64    `gorm:"type:bigint;comment:创建时间" json:"created_at"`
	UpdatedAt      int64    `gorm:"type:bigint;comment:更新时间" json:"updated_at"`
}
