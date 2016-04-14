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
	user := beego.NewNamespace("/user",
		beego.NSRouter("/list", &controllers.UserController{},"get:GetOne"),
		
	)
	beego.AddNamespace(user)
}
