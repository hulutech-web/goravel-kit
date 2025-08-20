package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/path"
	"goravel/packages/goravel-socket/servers"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("index.html")
	})

	facades.Route().Get("/crud", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("crud.html")
	})

	facades.Route().Static("/uploads", path.Public("uploads"))
	facades.Route().Static("/assets", path.Public("dist/assets"))

	websocketHandler := &servers.Controller{}
	facades.Route().Get("/ws", websocketHandler.Run)
	go servers.Manager.Start()
}
