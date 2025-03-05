package routes

import (
	"io/fs"
	"net/http"

	"goravel/public"

	"github.com/goravel/framework/facades"
)

func Web() {
	assetsCMS, _ := fs.Sub(&public.AssetsCMS, "cms")
	facades.Route().StaticFS("/cms", http.FS(assetsCMS))
	facades.Route().StaticFS("/static", http.Dir("./public/static"))
}
