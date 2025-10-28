package controller

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/path"
	httpfacade "github.com/hulutech-web/http_result"
	"goravel/packages/goravel_crud/core/controller"
	"goravel/packages/goravel_crud/core/curd_orm"
	"goravel/packages/goravel_crud/core/migration"
	"goravel/packages/goravel_crud/core/model"
	"goravel/packages/goravel_crud/core/request"
	"goravel/packages/goravel_crud/core/router"
	"goravel/packages/goravel_crud/validator"
	"strings"
)

type CRUDController struct {
	//Dependent services
}

func NewCRUDController() *CRUDController {
	return &CRUDController{
		//Inject services
	}
}

// 还要进行判断，当names时直接去掉s，如果是table_names则转换为TableName的形式
func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}

	// 处理复数形式（去掉末尾的's'）
	if len(s) > 1 && s[len(s)-1] == 's' {
		s = s[:len(s)-1]
	}

	// 分割下划线
	parts := strings.Split(s, "_")

	// 将每个部分首字母大写
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(string(part[0])) + part[1:]
		}
	}

	return strings.Join(parts, "")
}

func (r *CRUDController) Model(ctx http.Context) http.Response {
	name := capitalizeFirstLetter(ctx.Request().Input("name"))
	//首字母转大写

	err := model.Gen(name)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, err.Error(), nil)
	}
	return httpfacade.NewResult(ctx).Success("模型生成成功", nil)
}

func (r *CRUDController) Controller(ctx http.Context) http.Response {
	//去掉末尾的s
	tablename := ctx.Request().Input("tablename")
	if strings.HasSuffix(tablename, "s") {
		tablename = strings.TrimSuffix(tablename, "s")
	}
	name := capitalizeFirstLetter(tablename)
	err := controller.Gen(name)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, err.Error(), nil)
	}
	return httpfacade.NewResult(ctx).Success("控制器生成成功", nil)
}

func (r *CRUDController) Migration(ctx http.Context) http.Response {
	name := ctx.Request().Input("name")
	err := migration.Gen(name)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, err.Error(), nil)
	}
	return httpfacade.NewResult(ctx).Success("迁移文件生成成功", nil)
}

func (r *CRUDController) Request(ctx http.Context) http.Response {
	var ruleRequest validator.RuleRequest
	err := ctx.Request().Bind(&ruleRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, err.Error(), nil)
	}
	//去掉末尾的s
	if strings.HasSuffix(ruleRequest.Tablename, "s") {
		ruleRequest.Tablename = ruleRequest.Tablename[:len(ruleRequest.Tablename)-1]
	}
	name := capitalizeFirstLetter(ruleRequest.Tablename)
	err1 := request.Gen(name, ruleRequest.RuleFields)
	if err1 != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, err.Error(), nil)
	}
	return httpfacade.NewResult(ctx).Success("请求生成成功", nil)
}

func (r *CRUDController) Router(ctx http.Context) http.Response {
	tablename := ctx.Request().Input("tablename")
	if strings.HasSuffix(tablename, "s") {
		tablename = strings.TrimSuffix(tablename, "s")
	}
	name := capitalizeFirstLetter(tablename)
	err := router.Gen(name)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, err.Error(), nil)
	}
	return httpfacade.NewResult(ctx).Success("路由生成成功", nil)
}

//func (r *CRUDController) All(ctx http.Context) http.Response {
//	name := ctx.Request().Input("name")
//	upper_name := capitalizeFirstLetter(name)
//	err := model.Gen(upper_name)
//	if err != nil {
//		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, err.Error(), err.Error())
//	}
//
//	err = migration.Gen(name)
//	if err != nil {
//		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, err.Error(), err.Error())
//	}
//
//	err = request.Gen(upper_name)
//	if err != nil {
//		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, err.Error(), err.Error())
//	}
//
//	err = controller.Gen(upper_name)
//	if err != nil {
//		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, err.Error(), err.Error())
//	}
//
//	err = router.Gen(upper_name)
//	if err != nil {
//		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, err.Error(), err.Error())
//	}
//
//	return httpfacade.NewResult(ctx).Success("生成成功", nil)
//}

func (r *CRUDController) Tables(ctx http.Context) http.Response {
	tbs := curd_orm.TableDefine()
	return httpfacade.NewResult(ctx).Success("", tbs)
}

func (r *CRUDController) TableColumn(ctx http.Context) http.Response {
	tbname := ctx.Request().Input("table_name")
	tbs := curd_orm.TableSchema(tbname)
	return httpfacade.NewResult(ctx).Success("", tbs)
}

func (r *CRUDController) Migrate(ctx http.Context) http.Response {
	var genRequest validator.GenRequest
	err := ctx.Request().Bind(&genRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, err.Error(), nil)
	}
	sql := genRequest.Sql
	tablename := genRequest.Tablename
	//删除该表
	delTableSql := fmt.Sprintf("drop table %s", tablename)
	facades.Orm().Query().Exec(delTableSql)
	err1 := curd_orm.GenSql(sql)

	if err1 != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, err.Error(), nil)
	}
	path_model := path.App("models")
	err_file := curd_orm.GenModelFile(genRequest.Fields, tablename, path_model)
	if err_file != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, err_file.Error(), nil)
	}
	return httpfacade.NewResult(ctx).Success("建表成功", nil)
}
