package controller

func Gen(modelName string) error {
	//对modelName进行规范，当model首字母大写时不改动，当为UserRole时
	template := GenTemplate(modelName)
	err := CopyToCtrlPath(modelName, template)
	return err
}
