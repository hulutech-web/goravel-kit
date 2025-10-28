package factories

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/goravel/framework/facades"
)

type UserFactory struct {
}

// Definition Define the model's default state.
func (f *UserFactory) Definition() map[string]any {
	pwd, _ := facades.Hash().Make("admin888")
	return map[string]any{
		"username": gofakeit.Name(),
		"realname": gofakeit.Username(),
		"password": pwd,
	}
}
