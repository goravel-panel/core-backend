package permission

import (
	"goravel/app/http/controllers/admin"
	"goravel/app/services/permission"

	"github.com/goravel/framework/contracts/http"
)

type MenuController struct {
	response    admin.Response
	menuService permission.MenuService
}

func NewMenuController() *MenuController {
	return &MenuController{}
}

func (c *MenuController) All(ctx http.Context) http.Response {
	data, err := c.menuService.GetAll()
	if err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.Success(ctx, data)
}

func (c *MenuController) Detail(ctx http.Context) http.Response {
	id := ctx.Request().InputInt("id")
	menu, err := c.menuService.GetDetail(id)
	if err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.Success(ctx, menu)
}

func (c *MenuController) Save(ctx http.Context) http.Response {
	var request permission.MenuSaveRequest

	if err := ctx.Request().Bind(&request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	if err := c.menuService.Save(request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}

func (c *MenuController) Delete(ctx http.Context) http.Response {
	ids := ctx.Request().InputArray("id")
	if err := c.menuService.Delete(ids); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}
