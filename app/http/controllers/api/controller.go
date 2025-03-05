package api

import (
	"github.com/goravel/framework/contracts/http"
)

type Response struct {
}

// Success returns a success response with http code 200.
func (r Response) Success(ctx http.Context, data any, message ...string) http.Response {
	if len(message) == 0 {
		message = append(message, "")
	}
	return ctx.Response().Success().Json(http.Json{
		"code":    "Success",
		"message": message[0],
		"data":    data,
	})
}

// NoContent returns a no content response with http code 200.
func (r Response) NoContent(ctx http.Context) http.Response {
	return ctx.Response().NoContent(http.StatusOK)
}

// BadRequest returns a bad request response with http code 400.
func (r Response) BadRequest(ctx http.Context, message ...string) http.Response {
	if len(message) == 0 {
		message = append(message, http.StatusText(http.StatusBadRequest))
	}
	return ctx.Response().Status(http.StatusBadRequest).Json(http.Json{
		"code":    "BadRequest",
		"message": message[0],
		"data":    nil,
	})
}

// Unauthorized returns an unauthorized response with http code 401.
func (r Response) Unauthorized(ctx http.Context, message ...string) http.Response {
	if len(message) == 0 {
		message = []string{http.StatusText(http.StatusUnauthorized)}
	}
	return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{
		"code":    "Unauthorized",
		"message": message[0],
		"data":    nil,
	})
}

// Forbidden returns a forbidden response with http code 403.
func (r Response) Forbidden(ctx http.Context, message ...string) http.Response {
	if len(message) == 0 {
		message = []string{http.StatusText(http.StatusForbidden)}
	}
	return ctx.Response().Status(http.StatusForbidden).Json(http.Json{
		"code":    "Forbidden",
		"message": message[0],
		"data":    nil,
	})
}

// NotFound returns a not found response with http code 404.
func (Response) NotFound(ctx http.Context, message ...string) http.Response {
	if len(message) == 0 {
		message = []string{http.StatusText(http.StatusNotFound)}
	}
	return ctx.Response().Status(http.StatusNotFound).Json(http.Json{
		"code":    "NotFound",
		"message": message[0],
		"data":    nil,
	})
}

// ServerError returns a server error response with http code 500.
func (r Response) ServerError(ctx http.Context, message ...string) http.Response {
	if len(message) == 0 {
		message = []string{http.StatusText(http.StatusInternalServerError)}
	}
	return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{
		"code":    "ServerError",
		"message": message[0],
		"data":    nil,
	})
}
