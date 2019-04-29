// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"lovely_server/controllers"
	"lovely_server/models"
)

func init() {
	ns := beego.NewNamespace("/v1",
		// beego.NSNamespace("/api",
		// 	beego.NSInclude(
		// 		&controllers.ApiController{},
		// 	),
		// ),
		beego.NSNamespace("*",
			//Options用于跨域复杂请求预检
			beego.NSRouter("/*", &controllers.BaseController{}, "options:Options"),
		),
		beego.NSNamespace("/count",
			beego.NSInclude(
				&controllers.CountController{},
			),
		),
		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),
	)
	beego.Debug("models.Domain", models.Domain)
	beego.AddNamespace(ns)
	beego.InsertFilter(models.Domain, beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
}
