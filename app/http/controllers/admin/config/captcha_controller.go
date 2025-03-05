package config

import (
	"goravel/app/http/controllers/admin"
	"goravel/app/services/config"

	"github.com/goravel/framework/contracts/http"
)

type CaptchaController struct {
	response      admin.Response
	configService config.CaptchaService
}

func NewCaptchaController() *CaptchaController {
	return &CaptchaController{}
}

func (c *CaptchaController) Detail(ctx http.Context) http.Response {
	cfg, err := c.configService.GetConfig()
	if err != nil {
		return c.response.NotFound(ctx)
	}
	return c.response.Success(ctx, cfg)
}

func (c *CaptchaController) Save(ctx http.Context) http.Response {
	var request config.CaptchaConfig
	if err := ctx.Request().Bind(&request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	if err := c.configService.SetConfig(request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}
