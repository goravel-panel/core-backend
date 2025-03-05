package article

import (
	"goravel/app/http/controllers/admin"
	"goravel/app/models"
	"goravel/app/services"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type CategoryController struct {
	response admin.Response
}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

func (c *CategoryController) All(ctx http.Context) http.Response {
	var categories []models.ArticleCategory

	err := facades.Orm().Query().Order("sort desc,id desc").Get(&categories)
	if err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.Success(ctx, categories)
}

func (c *CategoryController) List(ctx http.Context) http.Response {
	pageNo := ctx.Request().QueryInt("pageNo", services.PageNo)
	pageSize := ctx.Request().QueryInt("pageSize", services.PageSize)
	name := ctx.Request().Query("name")
	isEnable := ctx.Request().QueryInt("isEnable", -1)

	var categories []models.ArticleCategory
	var total int64

	query := facades.Orm().Query()

	if len(name) > 0 {
		query = query.Where("name like ?", "%"+name+"%")
	}

	if isEnable > 0 {
		query = query.Where("is_enable", isEnable)
	}

	if err := query.Order("id desc").Paginate(pageNo, pageSize, &categories, &total); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}

	return c.response.Success(ctx, services.Paginate{
		Total:    total,
		PageNo:   pageNo,
		PageSize: pageSize,
		List:     categories,
	})
}

func (c *CategoryController) Detail(ctx http.Context) http.Response {
	id := ctx.Request().InputInt("id")

	var category models.ArticleCategory
	if err := facades.Orm().Query().FindOrFail(&category, id); err != nil {
		return c.response.NotFound(ctx)
	}
	return c.response.Success(ctx, category)
}
