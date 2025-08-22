const g=(r,e)=>{let t=`package models

import (
	"github.com/goravel/framework/database/orm"
	"goravel/packages/goravel-crud/core/curd_orm"
)

type ${p(r)} struct {
	orm.Model`;return e.forEach(o=>{const a=s(o.name),c=o.field_type,l=`gorm:"${o.gorm}" form:"${o.form}" json:"${o.json}"`;t+=`
	${a} ${c} \`${l}\``}),t+=`
}`,t+=`
    func init() {
   curd_orm.RegisterModel(&User{})
}`,t};function s(r){return r.replace(/([-_][a-z])/gi,e=>e.toUpperCase().replace("-","").replace("_",""))}function p(r){const e=m(r);return e.charAt(0).toUpperCase()+e.slice(1)}function m(r){return r.toLowerCase().replace(/([-_][a-z])/gi,e=>e.toUpperCase().replace("-","").replace("_","")).replace(/^./,e=>e.toLowerCase())}export{g as createModel};
//# sourceMappingURL=model.9a27e3f0.1755837291293.js.map
