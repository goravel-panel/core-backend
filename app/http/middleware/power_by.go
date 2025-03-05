package middleware

import (
	"fmt"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func PowerBy() http.Middleware {
	return func(ctx http.Context) {
		name := facades.Config().GetString("app.name", "gocmf")
		version := facades.Config().GetString("app.version")
		ctx.Response().Header("X-Powered-By", fmt.Sprintf("%s %s", name, version))
		ctx.Request().Next()
	}
}
