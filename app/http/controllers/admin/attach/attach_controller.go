package attach

import (
	"goravel/app/http/controllers/admin"
	"goravel/app/services"
	"goravel/app/services/attach"

	"github.com/goravel/framework/contracts/http"
)

type AttachController struct {
	response      admin.Response
	attachService attach.AttachService
}

func NewAttachController() *AttachController {
	return &AttachController{}
}

func (c *AttachController) List(ctx http.Context) http.Response {
	var request attach.AttachSearchRequest
	request.PageNo = ctx.Request().QueryInt("pageNo", services.PageNo)
	request.PageSize = ctx.Request().QueryInt("pageSize", services.PageSize)
	request.Name = ctx.Request().Query("name")
	request.Type = ctx.Request().Query("type")
	request.Mime = ctx.Request().Query("mime")
	request.CategoryID = ctx.Request().QueryInt("cid")

	data, err := c.attachService.GetList(request)
	if err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.Success(ctx, data)
}

func (c *AttachController) Rename(ctx http.Context) http.Response {
	id := ctx.Request().InputInt("id", 0)
	name := ctx.Request().Input("name")
	if err := c.attachService.Rename(id, name); err != nil {
		return c.response.BadRequest(ctx, "重命名失败："+err.Error())
	}
	return c.response.NoContent(ctx)
}

func (c *AttachController) Move(ctx http.Context) http.Response {
	cid := ctx.Request().InputInt("cid")
	ids := ctx.Request().InputArray("ids")
	if err := c.attachService.Move(ids, cid); err != nil {
		return c.response.BadRequest(ctx, "移动失败："+err.Error())
	}
	return c.response.NoContent(ctx)
}

func (c *AttachController) Delete(ctx http.Context) http.Response {
	ids := ctx.Request().InputArray("ids")
	if err := c.attachService.Delete(ids); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}
