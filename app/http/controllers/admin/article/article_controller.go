package article

import (
	"goravel/app/http/controllers/admin"
	"goravel/app/models"
	"goravel/app/services"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type ArticleController struct {
	response admin.Response
}

func NewArticleController() *ArticleController {
	return &ArticleController{}
}

func (c *ArticleController) List(ctx http.Context) http.Response {
	pageNo := ctx.Request().QueryInt("pageNo", services.PageNo)
	pageSize := ctx.Request().QueryInt("pageSize", services.PageSize)
	title := ctx.Request().Query("title")
	from := ctx.Request().Query("from")
	isEnable := ctx.Request().QueryInt("isEnable", -1)

	var articles []models.Article
	var total int64

	query := facades.Orm().Query().With("Category")

	if len(title) > 0 {
		query = query.Where("title like ?", "%"+title+"%")
	}

	if len(from) > 0 {
		query = query.Where("from", from)
	}

	if isEnable > 0 {
		query = query.Where("is_enable", isEnable)
	}

	if err := query.Order("id desc").Paginate(pageNo, pageSize, &articles, &total); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}

	return c.response.Success(ctx, services.Paginate{
		Total:    total,
		PageNo:   pageNo,
		PageSize: pageSize,
		List:     articles,
	})
}

func (c *ArticleController) Detail(ctx http.Context) http.Response {
	id := ctx.Request().InputInt("id")

	var article models.Article
	if err := facades.Orm().Query().FindOrFail(&article, id); err != nil {
		return c.response.NotFound(ctx, "文章不存在")
	}
	return c.response.Success(ctx, article.Resource())
}

func (c *ArticleController) Save(ctx http.Context) http.Response {
	type saveRequest struct {
		ID         int    `gorm:"primaryKey" json:"id"`
		CategoryID int    `json:"categoryID"`
		From       string `json:"from"`
		Title      string `json:"title"`
		Cover      string `json:"cover"`
		Content    string `json:"content"`
		IsEnable   int    `json:"isEnable"`
	}
	var request saveRequest

	if err := ctx.Request().Bind(&request); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}

	var article models.Article

	var where = map[string]any{}
	var data = map[string]any{}

	where["id"] = request.ID

	if len(request.Title) > 0 {
		data["title"] = request.Title
	}

	if len(request.Content) > 0 {
		data["content"] = request.Content
	}

	if len(request.From) > 0 {
		data["from"] = request.From
	}

	data["cover"] = request.Cover
	data["is_enable"] = request.IsEnable
	data["category_id"] = request.CategoryID

	err := facades.Orm().Query().UpdateOrCreate(&article, where, data)
	if err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}
	return c.response.NoContent(ctx)
}

func (c *ArticleController) Delete(ctx http.Context) http.Response {
	ids := ctx.Request().InputArray("id")

	if len(ids) < 0 {
		return c.response.BadRequest(ctx, "请选择要删除的数据")
	}
	var article models.Article
	if _, err := facades.Orm().Query().Where("id", ids).Delete(&article); err != nil {
		return c.response.BadRequest(ctx, "删除失败："+err.Error())
	}
	return c.response.NoContent(ctx)
}
