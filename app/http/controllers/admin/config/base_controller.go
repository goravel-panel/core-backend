package config

import (
	"goravel/app/http/controllers/admin"
	"goravel/app/services/config"

	"github.com/goravel/framework/contracts/http"
)

type BaseController struct {
	response      admin.Response
	configService config.BaseService
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (c *BaseController) Detail(ctx http.Context) http.Response {
	cfg, err := c.configService.GetConfig()
	if err != nil {
		return c.response.NotFound(ctx)
	}
	return c.response.Success(ctx, cfg)
}

func (c *BaseController) Save(ctx http.Context) http.Response {
	var request config.BaseConfig
	if err := ctx.Request().Bind(&request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	if err := c.configService.SetConfig(request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}
