package middleware

import (
	"errors"

	"goravel/app/models"

	"github.com/goravel/framework/auth"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func ApiJwt() http.Middleware {
	return func(ctx http.Context) {
		token := ctx.Request().Header("Authorization", "")
		if token == "" {
			ctx.Request().Abort(http.StatusUnauthorized)
			return
		}

		guard := facades.Auth(ctx).Guard("user")
		if _, err := guard.Parse(token); err != nil {
			if errors.Is(err, auth.ErrorTokenExpired) {
				token, err = guard.Refresh()
				if err != nil {
					// Refresh time exceeded
					ctx.Request().Abort(http.StatusUnauthorized)
				}
				token = "Bearer " + token
			} else {
				ctx.Request().Abort(http.StatusUnauthorized)
			}
		}

		var user models.User
		if err := facades.Auth(ctx).User(&user); err != nil {
			ctx.Request().Abort(http.StatusUnauthorized)
		}
		ctx.WithValue("user", user)
		ctx.Response().Header("Authorization", token)
		ctx.Request().Next()
	}

}
