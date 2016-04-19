package controllers

import (
	"mongoapi/modules"
	"github.com/astaxie/beego"
//	"log"
)

// Operations about object
type UserController struct {
	beego.Controller
}

// @Title GetOneById
// @Description find user by objectid
// @Param	objectId	"the objectid you want to get"
// @Success 200 {user} models.User
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (u *UserController) GetOneById() {
	objectId := u.Ctx.Input.Param(":objectId")
	if objectId != "" {
		ob, err := modules.GetOneById(objectId)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = ob
		}
	}
	u.ServeJSON()
}

// @Title GetOneByName
// @Description find user by filters
// @Param	objectId	"the objectid you want to get"
// @Success 200 {user} models.User
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (u *UserController) GetOneByName() {
	name :=  u.Ctx.Input.Param(":name")
	if name != "" {
		filters := map[string]string { "name" : name }
		ob, err := modules.GetOneByFilter(filters)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = ob
		}
	}
	u.ServeJSON()
}

// @Title GetAll
// @Description find user by objectid
// @Param   objectId  "the objectid you want to get"
// @Success 200 {user} models.User
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (u *UserController) GetAll() {
	ob, err := modules.GetAll()
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = ob
		}
	u.ServeJSON()
}
