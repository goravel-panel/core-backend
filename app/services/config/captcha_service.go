package config

import (
	"encoding/json"
	"errors"

	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type CaptchaConfig struct {
	Type          string `json:"type"`
	Length        int    `json:"length"`
	NoiseCount    int    `json:"noiseCount"`
	DigitSource   string `json:"digitSource"`
	LetterSource  string `json:"letterSource"`
	StringSource  string `json:"stringSource"`
	ChineseSource string `json:"chineseSource"`
}

type CaptchaService struct {
}

// GetConfig 获取验证码配置信息
func (s CaptchaService) GetConfig() (c CaptchaConfig, err error) {
	var cfg models.Config
	err = facades.Orm().Query().Where("type", models.TypeCaptcha).FirstOrFail(&cfg)
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

// SetConfig 设置验证码配置信息
func (s CaptchaService) SetConfig(c CaptchaConfig) (err error) {
	var cfg models.Config
	err = facades.Orm().Query().Where("type", models.TypeCaptcha).FirstOrFail(&cfg)
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
