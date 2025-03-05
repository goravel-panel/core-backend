package providers

import (
	"goravel/app/http"
	"goravel/routes"

	contracstHttp "github.com/goravel/framework/contracts/http"

	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/http/limit"
)

type RouteServiceProvider struct {
}

func (receiver *RouteServiceProvider) Register(app foundation.Application) {
}

func (receiver *RouteServiceProvider) Boot(app foundation.Application) {
	// Add HTTP middleware
	facades.Route().GlobalMiddleware(http.Kernel{}.Middleware()...)

	// Recover HTTP error
	facades.Route().Recover(func(ctx contracstHttp.Context, err any) {
		ctx.Request().Abort(contracstHttp.StatusInternalServerError)
	})

	// Configure rate limiting
	receiver.configureRateLimiting()

	routes.Web()
	routes.Admin()
	routes.Api()
}

func (receiver *RouteServiceProvider) configureRateLimiting() {
	facades.RateLimiter().For("sendMail", func(ctx contracstHttp.Context) contracstHttp.Limit {
		return limit.PerMinute(10).By(ctx.Request().Ip())
	})
	facades.RateLimiter().ForWithLimits("login", func(ctx contracstHttp.Context) []contracstHttp.Limit {
		return []contracstHttp.Limit{
			limit.PerDay(100),
			limit.PerMinute(5).By(ctx.Request().Ip()),
		}
	})
}
