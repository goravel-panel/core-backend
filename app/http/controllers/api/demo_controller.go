package api

import (
	"github.com/goravel/framework/contracts/http"
)

type DemoController struct {
	response Response
	// Dependent services
}

func NewDemoController() *DemoController {
	return &DemoController{
		// Inject services
	}
}

func (c *DemoController) Index(ctx http.Context) http.Response {
	return c.response.Success(ctx, http.Json{
		"message": "demo",
	})
}
