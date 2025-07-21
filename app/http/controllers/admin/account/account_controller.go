package account

import (
	"fmt"
	"goravel/app/models"

	adminController "goravel/app/http/controllers/admin"
	adminService "goravel/app/services/admin"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type AccountController struct {
	response       adminController.Response
	accountService adminService.AccountService
}

func NewAccountController() *AccountController {
	return &AccountController{}
}

// Login 用户登录
func (c *AccountController) Login(ctx http.Context) http.Response {
	var request adminService.LoginRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	token, err := c.accountService.Login(ctx, request)
	fmt.Println("token", token)
	if err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.Success(ctx, http.Json{
		"token": token,
	})
}

// Logout 退出登录
func (c *AccountController) Logout(ctx http.Context) http.Response {
	if err := facades.Auth(ctx).Guard(adminService.Guard).Logout(); err != nil {
		return c.response.BadRequest(ctx, "退出登录失败")
	}
	return c.response.NoContent(ctx)
}

// Profile 用户信息
func (c *AccountController) Profile(ctx http.Context) http.Response {
	admin, err := c.accountService.User(ctx)
	if err != nil {
		return c.response.Unauthorized(ctx, "登录状态失效，请重新登录")
	}
	menus, perms, err := c.accountService.GetPermissions(admin)
	if err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}

	type Data struct {
		Account string            `json:"account"`
		Name    string            `json:"name"`
		Avatar  string            `json:"avatar"`
		Menus   []models.TreeMenu `json:"menus"`
		Perms   []string          `json:"perms"`
	}
	return c.response.Success(ctx, &Data{
		Account: admin.Account,
		Name:    admin.Name,
		Avatar:  admin.Avatar,
		Menus:   menus,
		Perms:   perms,
	})
}

// Save 保存用户信息
func (c *AccountController) Save(ctx http.Context) http.Response {
	var request adminService.SaveRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	admin, err := c.accountService.User(ctx)
	if err != nil {
		return c.response.Unauthorized(ctx, "登录状态失效，请重新登录")
	}
	if err = c.accountService.Save(admin, request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	if len(request.NewPassword) > 0 {
		return c.response.Success(ctx, nil, "个人资料更新成功，重新登录")
	}
	return c.response.Success(ctx, nil, "个人资料更新成功")
}
