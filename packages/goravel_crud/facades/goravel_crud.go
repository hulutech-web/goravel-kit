package facades

import (
	"log"

	"goravel/packages/goravel_crud"
	"goravel/packages/goravel_crud/contracts"
)

func GoravelCrud() contracts.GoravelCrud {
	instance, err := goravel_crud.App.Make(goravel_crud.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.GoravelCrud)
}
