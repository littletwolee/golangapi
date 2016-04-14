package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["mongoapi/controllers:UserController"] = append(beego.GlobalControllerRouter["mongoapi/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

}
