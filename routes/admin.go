package routes

import (
	"goravel/app/http/controllers/admin/account"
	"goravel/app/http/controllers/admin/article"
	"goravel/app/http/controllers/admin/attach"
	"goravel/app/http/controllers/admin/captcha"
	"goravel/app/http/controllers/admin/config"
	"goravel/app/http/controllers/admin/home"
	"goravel/app/http/controllers/admin/permission"
	"goravel/app/http/controllers/admin/upload"
	"goravel/app/http/middleware"

	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	httpMiddleware "github.com/goravel/framework/http/middleware"
)

func Admin() {
	facades.Route().Prefix("backend").Group(func(r route.Router) {
		uploadController := upload.NewUploadController()
		captchaController := captcha.NewCaptchaController()

		accountController := account.NewAccountController()

		workbenchController := home.NewWorkbenchController()

		attachController := attach.NewAttachController()
		attachCategoryController := attach.NewCategoryController()

		permissionAdminController := permission.NewAdminController()
		permissionRoleController := permission.NewRoleController()
		permissionMenuController := permission.NewMenuController()

		configBaseController := config.NewBaseController()
		configCaptchaController := config.NewCaptchaController()

		// 公开接口
		r.Group(func(r route.Router) {
			// 配置信息
			r.Get("/config", configBaseController.Detail)
			// 账号登录
			r.Middleware(httpMiddleware.Throttle("login")).Post("/account/login", accountController.Login)
			// 验证码生成
			r.Post("/captcha/generate", captchaController.Generate)
			// 验证码验证
			r.Post("/captcha/verify", captchaController.Verify)
		})

		// 授权接口
		r.Middleware(middleware.AdminJwt()).Group(func(r route.Router) {
			// 用户资料
			r.Get("/account/profile", accountController.Profile)
			// 更新资料
			r.Post("/account/profile", accountController.Save)
			// 退出登录
			r.Get("/account/logout", accountController.Logout)

			// 仪表盘
			r.Get("/workbench", workbenchController.Home)
			// 图片上传
			r.Post("/upload/image", uploadController.Image)

			// 基础设置
			r.Get("/config/base/detail", configBaseController.Detail)
			r.Post("/config/base/save", configBaseController.Save)
			// 验证码设置
			r.Get("/config/captcha/detail", configCaptchaController.Detail)
			r.Post("/config/captcha/save", configCaptchaController.Save)

			// 附件分类
			r.Get("/attach/category/all", attachCategoryController.All)
			r.Post("/attach/category/save", attachCategoryController.Save)
			r.Post("/attach/category/rename", attachCategoryController.Rename)
			r.Post("/attach/category/delete", attachCategoryController.Delete)
			// 附件相关
			r.Get("/attach/list", attachController.List)
			r.Post("/attach/rename", attachController.Rename)
			r.Post("/attach/move", attachController.Move)
			r.Post("/attach/delete", attachController.Delete)

			// 权限用户
			r.Get("/permission/admin/list", permissionAdminController.List)
			r.Get("/permission/admin/detail", permissionAdminController.Detail)
			r.Post("/permission/admin/status", permissionAdminController.Status)
			r.Post("/permission/admin/save", permissionAdminController.Save)
			r.Post("/permission/admin/delete", permissionAdminController.Delete)
			// 权限角色
			r.Get("/permission/role/list", permissionRoleController.List)
			r.Get("/permission/role/all", permissionRoleController.All)
			r.Get("/permission/role/detail", permissionRoleController.Detail)
			r.Post("/permission/role/save", permissionRoleController.Save)
			r.Post("/permission/role/auth", permissionRoleController.Auth)
			r.Post("/permission/role/status", permissionRoleController.Status)
			r.Post("/permission/role/delete", permissionRoleController.Delete)
			// 权限菜单
			r.Get("/permission/menu/all", permissionMenuController.All)
			r.Get("/permission/menu/detail", permissionMenuController.Detail)
			r.Post("/permission/menu/save", permissionMenuController.Save)
			r.Post("/permission/menu/delete", permissionMenuController.Delete)
		})
	})

	// 业务接口
	facades.Route().Prefix("backend").Middleware(middleware.AdminJwt()).Group(func(r route.Router) {
		articleController := article.NewArticleController()
		articleCategoryController := article.NewCategoryController()

		r.Get("/article/list", articleController.List)
		r.Get("/article/detail", articleController.Detail)
		r.Post("/article/save", articleController.Save)
		r.Post("/article/delete", articleController.Delete)

		r.Get("/article/category/all", articleCategoryController.All)
		r.Get("/article/category/list", articleCategoryController.List)
		r.Get("/article/category/detail", articleCategoryController.Detail)
	})
}
