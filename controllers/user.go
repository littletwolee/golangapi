package controllers

import (
	"mongoapi/modules"
//	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about object
type UserController struct {
	beego.Controller
}

// @Title GetOne
// @Description find user by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {user} models.User
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *UserController) GetOne() {
	objectId := o.GetString("objectId")
	if objectId != "" {
		ob := modules.GetOne(objectId)
/*		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}*/
		o.Data["json"] = ob
	}
	
	o.ServeJSON()
}
