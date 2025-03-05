package permission

import (
	"goravel/app/http/controllers/admin"
	"goravel/app/services"
	"goravel/app/services/permission"

	"github.com/goravel/framework/contracts/http"
)

type RoleController struct {
	response    admin.Response
	roleService permission.RoleService
}

func NewRoleController() *RoleController {
	return &RoleController{}
}

func (c *RoleController) List(ctx http.Context) http.Response {
	var request permission.RoleSearchRequest

	request.PageNo = ctx.Request().QueryInt("pageNo", services.PageNo)
	request.PageSize = ctx.Request().QueryInt("pageSize", services.PageSize)
	request.Name = ctx.Request().Query("name")
	request.Remark = ctx.Request().Query("remark")
	request.IsEnable = ctx.Request().QueryInt("isEnable", 0)

	list, err := c.roleService.GetList(request)
	if err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.Success(ctx, list)
}

func (c *RoleController) All(ctx http.Context) http.Response {
	data, err := c.roleService.GetAll()
	if err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.Success(ctx, data)
}

func (c *RoleController) Detail(ctx http.Context) http.Response {
	id := ctx.Request().InputInt("id")
	role, err := c.roleService.GetDetail(id)
	if err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.Success(ctx, role)
}

func (c *RoleController) Save(ctx http.Context) http.Response {
	var request permission.RoleSaveRequest

	if err := ctx.Request().Bind(&request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	if err := c.roleService.Save(request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}

func (c *RoleController) Auth(ctx http.Context) http.Response {
	id := ctx.Request().InputInt("id", 0)
	menuIds := ctx.Request().InputArray("menuIds")

	role, err := c.roleService.GetDetail(id)
	if err != nil {
		c.response.NotFound(ctx)
	}
	if err = c.roleService.SetAuth(role, menuIds); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}

func (c *RoleController) Status(ctx http.Context) http.Response {
	id := ctx.Request().InputInt("id")
	isEnable := ctx.Request().InputInt("isEnable")

	role, err := c.roleService.GetDetail(id)
	if err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	if err = c.roleService.SetStatus(role, isEnable); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}

func (c *RoleController) Delete(ctx http.Context) http.Response {
	id := ctx.Request().InputInt("id")
	if err := c.roleService.Delete(id); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}
