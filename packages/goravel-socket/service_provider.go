package socket

import (
	"github.com/goravel/framework/contracts/foundation"
	"goravel/packages/goravel-socket/routers"
	"goravel/packages/goravel-socket/servers"
)

const Binding = "socket"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return nil, nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	routers.Init()
	go servers.WriteMessage()
	app.Publishes("goravel/packages/goravel-socket", map[string]string{
		"README.md": app.ConfigPath("socket.md"),
	})
}
