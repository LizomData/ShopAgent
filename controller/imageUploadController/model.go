package imageUploadController

type ImageUploadConfig struct {
	UploadDir    string   // 文件存储目录 (默认: ./tmp/image_upload)
	PublicPath   string   // 公开访问路径 (默认: /image/)
	MaxFileSize  int64    // 最大文件大小 (默认: 32MB)
	AllowedTypes []string // 允许的文件类型 (默认: 全部)
}

type ImageUploader struct {
	Config ImageUploadConfig
}

type QueryRequest struct {
	Page *int `json:"page" binding:"required"`
	Size *int `json:"size" binding:"required"`
}
