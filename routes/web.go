package routes

import (
	"net/http"

	"github.com/goravel/framework/facades"
)

func Web() {
	facades.Route().StaticFS("/static", http.Dir("./public/static"))
}
