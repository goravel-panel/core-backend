package config

import (
	"encoding/json"
	"errors"

	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type BaseConfig struct {
	Name            string `json:"name"`
	Logo            string `json:"logo"`
	Email           string `json:"email"`
	Copyright       string `json:"copyright"`
	LoginBackground string `json:"loginBackground"`
	UseLoginCaptcha bool   `json:"useLoginCaptcha"`
}

type BaseService struct {
}

// GetConfig 获取基础配置信息
func (s BaseService) GetConfig() (c BaseConfig, err error) {
	var cfg models.Config
	err = facades.Orm().Query().Where("type = ?", models.TypeBase).FirstOrFail(&cfg)
	if err != nil {
		err = errors.New("配置信息不存在")
		return
	}
	if len(cfg.Value) == 0 {
		return
	}
	err = json.Unmarshal([]byte(cfg.Value), &c)
	return
}

// SetConfig 设置基础配置信息
func (s BaseService) SetConfig(c BaseConfig) (err error) {
	var cfg models.Config
	err = facades.Orm().Query().Where("type = ?", models.TypeBase).FirstOrFail(&cfg)
	if err != nil {
		err = errors.New("配置信息不存在")
		return err
	}
	value, err := json.Marshal(c)
	if err != nil {
		return err
	}
	_, err = facades.Orm().Query().Model(&cfg).Update("value", string(value))
	return err
}
