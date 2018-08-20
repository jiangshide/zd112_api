// @APIVersion 1.0.0
// @Title zd112 API
// @Description this for mobile with zd112
// @Contact 18311271399@163.com
// @TermsOfServiceUrl http://www.zd112.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"zd112_api/controllers"
	"github.com/astaxie/beego"
)

func init() {

	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/home",
			beego.NSInclude(
				&controllers.HomeController{},
			),
		),
		beego.NSNamespace("/upload",
			beego.NSInclude(
				&controllers.UploadController{},
			), ),
		beego.NSNamespace("/blockchain",
			beego.NSInclude(
				&controllers.BlockChainController{},
			),
		),
		beego.NSNamespace("/zdb",
			beego.NSInclude(
				&controllers.AppController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.ErrorController(&controllers.ErrorController{})
}
