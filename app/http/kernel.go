package http

import (
	"goravel/app/http/middleware"

	"github.com/goravel/framework/contracts/http"
)

type Kernel struct {
}

// Middleware The application's global HTTP middleware stack.
// These middleware are run during every request to your application.
func (kernel Kernel) Middleware() []http.Middleware {
	return []http.Middleware{
		middleware.PowerBy(),
	}
}
