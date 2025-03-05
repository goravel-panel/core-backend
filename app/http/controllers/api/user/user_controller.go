package user

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/http/controllers/api"
	"goravel/app/models"
)

type UserController struct {
	response api.Response
}

func NewUserController() *UserController {
	return &UserController{
		// Inject services
	}
}

func (c *UserController) Show(ctx http.Context) http.Response {
	var user models.User
	if err := facades.Orm().Query().FindOrFail(&user, 1); err != nil {
		return c.response.BadRequest(ctx, err.Error())
	}

	return c.response.Success(ctx, user)
}
