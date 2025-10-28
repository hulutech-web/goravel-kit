package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
)

type AppServiceProvider struct {
}

func (receiver *AppServiceProvider) Register(app foundation.Application) {

}

func (receiver *AppServiceProvider) Boot(app foundation.Application) {
	facades.View().Share("pdf_save_path", facades.Config().Get("pdf_save_path"))
	facades.View().Share("pdf_prefix", facades.Config().Get("pdf_prefix"))
	facades.View().Share("getDefaultTemplate", facades.Config().Get("getDefaultTemplate"))
	facades.View().Share("saveTemplate", facades.Config().Get("saveTemplate"))
	facades.View().Share("saveHTML", facades.Config().Get("saveHTML"))
	facades.View().Share("getIndexTemplate", facades.Config().Get("getIndexTemplate"))
	facades.View().Share("generate", facades.Config().Get("generate"))
}
