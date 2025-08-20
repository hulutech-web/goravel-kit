package request

import "goravel/packages/goravel_crud/validator"

func Gen(modelName string, fields []validator.RuleField) error {
	template := GenTemplate(modelName, fields)
	err := CopyAndFillRuleToRequestPath(modelName, template)
	return err
}
