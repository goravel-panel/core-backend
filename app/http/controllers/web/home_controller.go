package web

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/support"
)

type HomeController struct {
	// Dependent services
}

func NewHomeController() *HomeController {
	return &HomeController{
		// Inject services
	}
}

func (r *HomeController) Show(ctx http.Context) http.Response {
	return ctx.Response().View().Make("welcome.tmpl", map[string]any{
		"version": support.Version,
	})
}
