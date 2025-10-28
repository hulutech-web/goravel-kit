package config

import (
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/path"
)

func init() {
	config := facades.Config()
	
	/*方法路径*/
	config.Add("pdf_save_path", path.Public("pdf"))
	config.Add("pdf_prefix", "pdf_prefix")
	config.Add("getDefaultTemplate", "getDefaultTemplate")
	config.Add("saveTemplate", "saveTemplate")
	config.Add("saveHTML", "saveHTML")
	config.Add("getIndexTemplate", "getIndexTemplate")
	config.Add("generate", "generate")
}
