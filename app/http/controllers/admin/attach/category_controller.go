package attach

import (
	"goravel/app/http/controllers/admin"
	"goravel/app/services/attach"

	"github.com/goravel/framework/contracts/http"
)

type CategoryController struct {
	response        admin.Response
	categoryService attach.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

func (c *CategoryController) All(ctx http.Context) http.Response {
	data, err := c.categoryService.GetAll()
	if err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.Success(ctx, data)
}

func (c *CategoryController) Save(ctx http.Context) http.Response {
	var request attach.CategorySaveRequest
	request.ParentID = ctx.Request().InputInt("pid", 0)
	request.Name = ctx.Request().Input("name")

	if err := c.categoryService.Save(request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}

func (c *CategoryController) Rename(ctx http.Context) http.Response {
	id := ctx.Request().InputInt("id", 0)
	name := ctx.Request().Input("name")

	if err := c.categoryService.Rename(id, name); err != nil {
		return c.response.BadRequest(ctx, "重命名失败："+err.Error())
	}
	return c.response.NoContent(ctx)
}

func (c *CategoryController) Delete(ctx http.Context) http.Response {
	id := ctx.Request().InputInt("id")
	if err := c.categoryService.Delete(id); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}
