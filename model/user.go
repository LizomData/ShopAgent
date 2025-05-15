package model

type User struct {
	ID       int64  `gorm:"primary_key;type:bigint" json:"id"`
	Username string `gorm:"type:varchar(255);unique;not null" json:"username" binding:"required"`
	Password string `gorm:"type:varchar(255);not null" json:"password" binding:"required"`
}
