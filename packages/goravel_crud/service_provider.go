package goravel_crud

import (
	"github.com/goravel/framework/contracts/binding"
	"github.com/goravel/framework/contracts/foundation"
	"goravel/packages/goravel_crud/routes"
)

const Binding = "goravel_crud"

var App foundation.Application

type ServiceProvider struct {
}

// Relationship returns the relationship of the service provider.
func (r *ServiceProvider) Relationship() binding.Relationship {
	return binding.Relationship{
		Bindings:     []string{},
		Dependencies: []string{},
		ProvideFor:   []string{},
	}
}

// Register registers the service provider.
func (r *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return nil, nil
	})
}

// Boot boots the service provider, will be called after all service providers are registered.
func (r *ServiceProvider) Boot(app foundation.Application) {
	routes.CRUD(app)
	app.Publishes("./packages/goravel-crud", map[string]string{
		"config/crud.go":        app.ConfigPath("crud.go"),
		"routes/crud_api.go":    app.BasePath("routes"),
		"panel/dist":            app.PublicPath("panel"),
		"panel/dist/index.html": app.BasePath("resources/views/index.html"),
	})
}
