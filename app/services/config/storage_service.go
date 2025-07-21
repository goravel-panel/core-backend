package config

import (
	"encoding/json"
	"errors"

	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type StorageConfig struct {
	Name      string `json:"name"`
	Bucket    string `json:"bucket"`
	SecretId  string `json:"secretId"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Endpoint  string `json:"endpoint"`
	Domain    string `json:"domain"`
}

type StorageService struct {
}

// GetAll 获取所有存储驱动
func (s StorageService) GetAll() (configs []models.Config, err error) {
	err = facades.Orm().Query().Model(&configs).Where("type = ? and is_enable = ?", models.TypeStorage, models.IsEnable).With("Attaches").Order("sort desc,id desc").Get(&configs)
	return
}

// GetConfig 获取指定存储配置信息
func (s StorageService) GetConfig(name string) (c StorageConfig, err error) {
	var cfg models.Config
	err = facades.Orm().Query().Where("type = ? and name = ?", models.TypeStorage, name).FirstOrFail(&cfg)
	if err != nil {
		err = errors.New("配置信息不存在")
		return
	}
	if len(cfg.Value) == 0 {
		return
	}
	err = json.Unmarshal([]byte(cfg.Value), &c)
	c.Name = name
	return
}

// SetConfig 设置存储配置信息
func (s StorageService) SetConfig(c StorageConfig, isDefault int) error {
	var cfg models.Config
	err := facades.Orm().Query().Where("type = ? and name =?", models.TypeStorage, c.Name).FirstOrFail(&cfg)
	if err != nil {
		return errors.New("配置信息不存在")
	}
	if isDefault == -models.IsDefault {
		total, _ := facades.Orm().Query().Model(&models.Config{}).Where("type = ? and is_default = ? and id <> ?", models.TypeStorage, models.IsDefault, cfg.ID).Count()
		if total == 0 {
			return errors.New("必须设置一个默认驱动")
		}
	}

	if isDefault == models.IsDefault {
		_, _ = facades.Orm().Query().Model(&models.Config{}).Where("type = ? and id <> ?", models.TypeStorage, cfg.ID).Update("is_default", -models.IsDefault)
	}

	value, err := json.Marshal(c)
	if err != nil {
		return err
	}
	cfg.IsDefault = isDefault
	cfg.Value = string(value)
	return facades.Orm().Query().Save(&cfg)
}

// GetDefault 获取默认存储配置信息
func (s StorageService) GetDefault() (c StorageConfig, err error) {
	var cfg models.Config
	err = facades.Orm().Query().Where("type = ? and is_default = ?", models.TypeStorage, models.IsDefault).FirstOrFail(&cfg)
	if err != nil {
		err = errors.New("配置信息不存在")
		return
	}
	if len(cfg.Value) == 0 {
		return
	}
	err = json.Unmarshal([]byte(cfg.Value), &c)
	c.Name = cfg.Name
	return
}

// SetDefault 设置为默认
func (s StorageService) SetDefault(name string, isDefault int) (err error) {
	var cfg models.Config
	err = facades.Orm().Query().Where("type = ? and name = ?", models.TypeStorage, name).FirstOrFail(&cfg)
	if err != nil {
		return errors.New("配置信息不存在")
	}
	if isDefault == -models.IsDefault {
		var total int64
		if total, err = facades.Orm().Query().Model(&cfg).Where("type = ? and is_default = ? and id <> ?", models.TypeStorage, models.IsDefault, cfg.ID).Count(); err != nil || total == 0 {
			return errors.New("必须设置一个默认驱动")
		}
	}
	if isDefault == models.IsDefault {
		_, err = facades.Orm().Query().Model(&models.Config{}).Where("type = ? and id <> ?", models.TypeStorage, cfg.ID).Update("is_default", -models.IsDefault)
		if err != nil {
			return err
		}
	}
	_, err = facades.Orm().Query().Model(&models.Config{}).Where("id", cfg.ID).Update("is_default", isDefault)
	return err
}
