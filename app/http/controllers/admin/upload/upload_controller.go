package upload

import (
	"goravel/app/http/controllers/admin"
	"goravel/app/services/upload"

	"github.com/goravel/framework/contracts/http"
)

type UploadController struct {
	service  upload.UploadService
	response admin.Response
}

func NewUploadController() *UploadController {
	return &UploadController{}
}

func (c *UploadController) Image(ctx http.Context) http.Response {
	categoryID := ctx.Request().InputInt("cid")

	file, err := ctx.Request().File("image")
	if err != nil {
		return c.response.BadRequest(ctx, "上传失败："+err.Error())
	}

	url, err := c.service.FromFile(file).Upload(categoryID, "images")
	if err != nil {
		return c.response.BadRequest(ctx, "上传失败："+err.Error())
	}
	return c.response.Success(ctx, url)
}
