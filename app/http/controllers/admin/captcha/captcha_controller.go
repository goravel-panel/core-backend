package captcha

import (
	"goravel/app/http/controllers/admin"
	"goravel/app/services/captcha"

	"github.com/goravel/framework/contracts/http"
)

type CaptchaController struct {
	response       admin.Response
	captchaService captcha.CaptchaService
}

func NewCaptchaController() *CaptchaController {
	return &CaptchaController{}
}

// Generate 生成验证码
func (c *CaptchaController) Generate(ctx http.Context) http.Response {
	var request captcha.GenerateRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	id, src, err := c.captchaService.Generate(request)
	if err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.Success(ctx, http.Json{
		"id":  id,
		"src": src,
	})
}

// Verify 校验验证码
func (c *CaptchaController) Verify(ctx http.Context) http.Response {
	var request captcha.VerifyRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	if err := c.captchaService.Verify(request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}
