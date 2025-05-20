package model

type ImageUploadRecord struct {
	ID        int64   `gorm:"type:bigint;primary_key;comment:记录id" json:"id,string"`
	UserID    int64   `gorm:"type:bigint;not null;comment:关联用户ID" json:"-"`
	PublicURL *string `gorm:"type:varchar(255);not null;default:'';comment:公共访问url" json:"public_url"`
	FileName  *string `gorm:"type:varchar(255);not null;default:'';comment:文件名字" json:"file_name"`
	FilePath  *string `gorm:"type:varchar(255);not null;default:'';comment:文件存储路径" json:"file_path"`
	FileSize  *int64  `gorm:"type:bigint;not null;default:0;comment:文件大小" json:"file_size"`
	CreatedAt int64   `gorm:"type:bigint;comment:创建时间" json:"created_at"`
	UpdatedAt int64   `gorm:"type:bigint;comment:更新时间" json:"updated_at"`
}
