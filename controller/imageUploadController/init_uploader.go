package imageUploadController

import (
	"ShopAgent/model"
	"ShopAgent/util"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

var Uploader = InitUploader()

func InitUploader() *ImageUploader {
	// 初始化上传模块
	_uploader, err := NewUploader(ImageUploadConfig{
		UploadDir:    "./tmp/image_upload",
		PublicPath:   "/image/",
		MaxFileSize:  100 << 20, // 50MB
		AllowedTypes: []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg"},
	})
	if err != nil {
		panic(fmt.Sprintf("初始化上传模块失败: %v", err))
	}

	return _uploader
}

// 初始化上传器
func NewUploader(cfg ImageUploadConfig) (*ImageUploader, error) {
	// 设置默认值
	if cfg.UploadDir == "" {
		cfg.UploadDir = "./tmp/image_upload"
	}
	if cfg.PublicPath == "" {
		cfg.PublicPath = "/image/"
	}
	if cfg.MaxFileSize == 0 {
		cfg.MaxFileSize = 32 << 20 // 32MB
	}

	// 创建存储目录
	if err := os.MkdirAll(cfg.UploadDir, 0755); err != nil {
		return nil, fmt.Errorf("创建上传目录失败: %w", err)
	}

	return &ImageUploader{Config: cfg}, nil
}

// 核心上传方法
func (u *ImageUploader) UploadFile(fileHeader *multipart.FileHeader) (*model.ImageUploadRecord, error) {
	// 验证文件大小
	if fileHeader.Size > u.Config.MaxFileSize {
		return nil, fmt.Errorf("文件超过大小限制(最大 %dMB)", u.Config.MaxFileSize/(1<<20))
	}

	// 验证文件类型
	if len(u.Config.AllowedTypes) > 0 {
		ext := filepath.Ext(fileHeader.Filename)
		if !u.contains(u.Config.AllowedTypes, ext) {
			return nil, fmt.Errorf("不支持的文件类型: %s", ext)
		}
	}

	// 生成唯一文件名
	newFilename := util.GenerateUniqueFilename(fileHeader.Filename)

	// 保存文件
	dstPath := filepath.Join(u.Config.UploadDir, newFilename)
	if err := u.SaveUploadedFile(fileHeader, dstPath); err != nil {
		return nil, fmt.Errorf("文件保存失败: %w", err)
	}

	// 生成访问地址
	publicURL := filepath.Join(u.Config.PublicPath, newFilename)

	return &model.ImageUploadRecord{
		PublicURL: &publicURL,
		FilePath:  &dstPath,
		FileName:  &newFilename,
		FileSize:  &fileHeader.Size,
	}, nil
}

// 辅助方法：检查切片包含
func (u *ImageUploader) contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// 保存方法实现
func (u *ImageUploader) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	// 1. 打开上传文件
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("无法打开上传文件: %w", err)
	}
	defer src.Close()

	// 2. 创建目标文件
	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("无法创建目标文件: %w", err)
	}
	defer out.Close()

	// 3. 复制文件内容
	if _, err := io.Copy(out, src); err != nil {
		return fmt.Errorf("文件写入失败: %w", err)
	}

	return nil
}
