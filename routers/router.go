// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"mongoapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/user",
			beego.NSRouter("/Id/?:objectId", &controllers.UserController{},"get:GetOneUserById"),
			beego.NSRouter("/Name/?:name", &controllers.UserController{},"get:GetOneUserByName"),
			beego.NSRouter("/list", &controllers.UserController{},"get:GetAllUsers"),
			beego.NSRouter("/Create", &controllers.UserController{},"post:CreateUser"),
			beego.NSRouter("/Delete/?:objectId", &controllers.UserController{},"delete:DeleteUser"),
		),		
	)
	beego.AddNamespace(ns)
}
