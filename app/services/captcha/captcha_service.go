package captcha

import (
	"errors"

	"goravel/app/services/config"

	"github.com/golang-module/base64Captcha"
	"github.com/golang-module/base64Captcha/driver"
	"github.com/golang-module/base64Captcha/store"
)

type CaptchaService struct {
	configService config.CaptchaService
}

var base64Driver driver.Driver
var base64Store = store.DefaultStoreMemory

type GenerateRequest struct {
	Type       string `form:"type"`
	Length     int    `form:"length"`
	Height     int    `form:"height"`
	Width      int    `form:"width"`
	NoiseCount int    `form:"noiseCount"`
}

func (s *CaptchaService) Generate(r GenerateRequest) (id, b64s string, err error) {
	cfg, err := s.configService.GetConfig()

	if err != nil {
		err = errors.New("请到系统设置->验证码设置里配置必要参数")
	}

	if len(r.Type) == 0 {
		r.Type = cfg.Type
	}

	if r.Length <= 0 {
		r.Length = cfg.Length
	}

	if r.Width <= 0 {
		r.Width = 100
	}

	if r.Height <= 0 {
		r.Height = 30
	}

	if r.NoiseCount <= 0 {
		r.NoiseCount = cfg.NoiseCount
	}

	driverConfig := struct {
		CaptchaType  string
		Fonts        []string
		Source       string
		DriverString driver.DriverString
		DriverLetter driver.DriverLetter
		DriverDigit  driver.DriverDigit
		DriverMath   driver.DriverMath
	}{
		CaptchaType: r.Type,
		DriverString: driver.DriverString{
			Width:      r.Width,
			Height:     r.Height,
			Length:     r.Length,
			NoiseCount: r.NoiseCount,
		},
		DriverLetter: driver.DriverLetter{
			Width:      r.Width,
			Height:     r.Height,
			Length:     r.Length,
			NoiseCount: r.NoiseCount,
		},
		DriverDigit: driver.DriverDigit{
			Width:      r.Width,
			Height:     r.Height,
			Length:     r.Length,
			NoiseCount: r.NoiseCount,
		},
		DriverMath: driver.DriverMath{
			Width:      r.Width,
			Height:     r.Height,
			NoiseCount: r.NoiseCount,
		},
	}

	switch driverConfig.CaptchaType {
	case "letter":
		base64Driver = driver.NewDriverLetter(driverConfig.DriverLetter)
	case "digit":
		base64Driver = driver.NewDriverDigit(driverConfig.DriverDigit)
	case "math":
		base64Driver = driver.NewDriverMath(driverConfig.DriverMath)
	default:
		base64Driver = driver.NewDriverString(driverConfig.DriverString)
	}

	captcha := base64Captcha.NewCaptcha(base64Driver, base64Store)
	id, b64s, _, err = captcha.Generate()
	if err != nil {
		err = errors.New("验证码生成失败：" + err.Error())
	}
	return
}

type VerifyRequest struct {
	ID    string `form:"id"`
	Value string `form:"value"`
}

func (s *CaptchaService) Verify(r VerifyRequest) error {
	if !base64Store.Verify(r.ID, r.Value, true) {
		return errors.New("验证码验证失败")
	}
	return nil
}
