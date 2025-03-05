package upload

import (
	"errors"
	"fmt"
	"path/filepath"

	"goravel/app/models"
	"goravel/app/services/config"

	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/carbon"
)

type UploadService struct {
	disk string
	dir  string
	file filesystem.File
	url  string

	// 依赖注入
	configService config.StorageService
}

// SetConfig 设置配置信息
func (s UploadService) SetConfig(disk string) error {
	cfg, err := s.configService.GetConfig(disk)
	if err != nil {
		return err
	}
	switch disk {
	case "local":
		return nil
	case "oss":
		facades.Config().Add("filesystems.disks.oss.key", cfg.AccessKey)
		facades.Config().Add("filesystems.disks.oss.secret", cfg.SecretKey)
		facades.Config().Add("filesystems.disks.oss.bucket", cfg.Bucket)
		facades.Config().Add("filesystems.disks.oss.url", cfg.Domain)
		facades.Config().Add("filesystems.disks.oss.endpoint", cfg.Endpoint)
	case "cos":
		facades.Config().Add("filesystems.disks.cos.key", cfg.SecretId)
		facades.Config().Add("filesystems.disks.cos.secret", cfg.SecretKey)
		facades.Config().Add("filesystems.disks.cos.bucket", cfg.Bucket)
		facades.Config().Add("filesystems.disks.cos.url", cfg.Domain)
	default:
		return errors.New("没有设置默认存储引擎，请到系统设置->存储设置里配置")
	}
	return nil
}

// FromFile 从本地文件上传
func (s UploadService) FromFile(file filesystem.File) UploadService {
	s.file = file
	return s
}

func (s UploadService) Upload(categoryID int, dir ...string) (url string, err error) {
	cfg, err := s.configService.GetDefault()
	if err != nil {
		return
	}
	disk := cfg.Name
	if len(cfg.Name) == 0 {
		disk = s.disk
	}

	if err = s.SetConfig(disk); err != nil {
		return
	}

	fileDir := ""
	if len(dir) > 0 {
		fileDir = fmt.Sprintf("%s%s%s", dir[0], string(filepath.Separator), carbon.Now().ToShortDateString())
	}

	var driver filesystem.Driver
	if disk == "local" {
		driver = facades.Storage().Disk("public")
	} else {
		driver = facades.Storage().Disk(disk)
	}
	filePath, err := driver.PutFile(fileDir, s.file)
	if err != nil {
		return
	}

	var attach models.Attach

	fileInfo := getFileInfo(s.file)

	attach.Type = models.TypeImage
	attach.CategoryID = categoryID
	attach.Disk = disk
	attach.Name = fileInfo.Name
	attach.Ext = fileInfo.Extension
	attach.Size = fileInfo.Size
	attach.Mime = fileInfo.Mime
	attach.Path = filePath

	attach.Url = driver.Url(filePath)
	err = facades.Orm().Query().Create(&attach)
	return
}

type FileInfo struct {
	Name      string
	Extension string
	Mime      string
	Size      int64
}

func getFileInfo(file filesystem.File) (f FileInfo) {
	f.Name = file.GetClientOriginalName()
	f.Size, _ = file.Size()
	f.Extension, _ = file.Extension()
	f.Mime, _ = file.MimeType()
	return
}
