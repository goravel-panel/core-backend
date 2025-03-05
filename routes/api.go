package routes

import (
	"goravel/app/http/controllers/api"
	"goravel/app/http/controllers/api/user"
	"goravel/app/http/middleware"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func Api() {
	facades.Route().Prefix("api").Middleware(middleware.ApiSign()).Group(func(r route.Router) {
		userController := user.NewUserController()
		demoController := api.NewDemoController()

		// 公开接口
		r.Group(func(r route.Router) {
			r.Get("/", func(ctx http.Context) http.Response {
				return ctx.Response().Success().String("恭喜您 API 服务启动成功~")
			})
			r.Get("user", userController.Show)
			r.Get("demo", demoController.Index)
		})

		// 授权接口
		r.Middleware(middleware.ApiJwt()).Group(func(r route.Router) {

		})
	})
}
