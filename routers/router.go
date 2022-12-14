// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"my-web-test/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			//beego.NSRouter("/", &controllers.UserController{}, "Post:Post"),
			beego.NSRouter("/", &controllers.UserController{}),
		),
		beego.NSNamespace("/yang"),//beego.NSRouter("/getone", &controllers.YangController{}, "GET:GetOne"),
		//beego.NSRouter("/get", &controllers.YangController{}),
		//beego.NSInclude(&controllers.YangController{}),

	)
	beego.AutoRouter(&controllers.YangController{})
	beego.AddNamespace(ns)

}
