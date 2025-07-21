package middleware

import (
	"errors"
	"fmt"

	"goravel/app/models"
	"goravel/app/services/admin"

	"github.com/goravel/framework/auth"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func AdminJwt() http.Middleware {
	return func(ctx http.Context) {
		token := ctx.Request().Header("Authorization", "")
		if token == "" {
			ctx.Request().Abort(http.StatusUnauthorized)
			return
		}

		guard := facades.Auth(ctx).Guard(admin.Guard)
		if _, err := guard.Parse(token); err != nil {
			fmt.Println("Parse err", err)
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

		var account models.Admin
		if err := guard.User(&account); err != nil {
			ctx.Request().Abort(http.StatusUnauthorized)
		}
		if account.IsEnable != models.IsEnable {
			ctx.Request().Abort(http.StatusForbidden)
		}
		ctx.WithValue(admin.Guard, account)
		ctx.Response().Header("Authorization", token)
		ctx.Request().Next()
	}

}
