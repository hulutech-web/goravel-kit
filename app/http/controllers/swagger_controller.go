package controllers

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "goravel/docs"
	_ "net/url"
)

type SwaggerController struct {
	//Dependent services
}

func NewSwaggerController() *SwaggerController {
	return &SwaggerController{
		//Inject services
	}
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host petstore.swagger.io
// @BasePath /api
func (r *SwaggerController) Index(ctx http.Context) http.Response {
	// 配置 Swagger UI
	handler := httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
		httpSwagger.BeforeScript(`const TokenInterceptorPlugin = () => ({
  fn: {
    // 拦截请求并添加 Authorization header
    addAuthHeader: (req) => {
      const token = localStorage.getItem('swagger_token');
      if (token) {
        req.headers['Authorization'] = 'Bearer ' + token;
      }
      return req;
    }
  },
  // requestInterceptor 会在每次请求前调用
  requestInterceptor: (req) => {
    return req.addAuthHeader(req);
  }
});`),
		// 注册自定义插件 TokenInterceptorPlugin
		httpSwagger.Plugins([]string{"TokenInterceptorPlugin"}),
		// 配置 UI 的行为，包括用户输入 Token 后保存到 localStorage
		httpSwagger.UIConfig(map[string]string{
			"onComplete": fmt.Sprintf(`() => {
    // 监听 API Key 输入并将 Token 存储在 localStorage
    const apiKeyInput = document.querySelector('#input_apiKey');
    if (apiKeyInput) {
      apiKeyInput.addEventListener('change', (e) => {
        const token = e.target.value;
        localStorage.setItem('swagger_token', token);
      });
    }
  }`),
		}),
	)

	// 返回 Swagger UI 页面
	handler(ctx.Response().Writer(), ctx.Request().Origin())
	return nil
}
