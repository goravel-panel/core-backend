package permission

import (
	"goravel/app/http/controllers/admin"
	"goravel/app/services"
	"goravel/app/services/permission"

	"github.com/goravel/framework/contracts/http"
)

type AdminController struct {
	response     admin.Response
	adminService permission.AdminService
}

func NewAdminController() *AdminController {
	return &AdminController{}
}

func (c *AdminController) List(ctx http.Context) http.Response {
	var request permission.AdminSearchRequest

	request.PageNo = ctx.Request().QueryInt("pageNo", services.PageNo)
	request.PageSize = ctx.Request().QueryInt("pageSize", services.PageSize)
	request.Name = ctx.Request().Query("name")
	request.Account = ctx.Request().Query("account")
	request.IsEnable = ctx.Request().QueryInt("isEnable", 0)
	request.RoleID = ctx.Request().QueryInt("roleId", 0)

	list, err := c.adminService.GetList(request)
	if err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.Success(ctx, list)
}

func (c *AdminController) Detail(ctx http.Context) http.Response {
	id := ctx.Request().InputInt("id")
	data, err := c.adminService.GetDetail(id)
	if err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.Success(ctx, data)
}

func (c *AdminController) Save(ctx http.Context) http.Response {
	var request permission.AdminSaveRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	if err := c.adminService.Save(request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}

func (c *AdminController) Status(ctx http.Context) http.Response {
	id := ctx.Request().InputInt("id")
	isEnable := ctx.Request().InputInt("isEnable")
	data, err := c.adminService.GetDetail(id)
	if err != nil {
		c.response.NotFound(ctx)
	}
	if err = c.adminService.SetStatus(data, isEnable); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}

func (c *AdminController) Delete(ctx http.Context) http.Response {
	id := ctx.Request().InputInt("id")
	if err := c.adminService.Delete(id); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}
