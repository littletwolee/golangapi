package controllers

import (
	"mongoapi/modules"
	"mongoapi/models"
	"github.com/astaxie/beego"
	"encoding/json"
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
		ob, err := (&modules.User{}).GetOneById(objectId)
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
// @Param	name	"the name you want to get"
// @Success 200 {user} models.User
// @Failure 403 :name is empty
// @router /:name [get]
func (u *UserController) GetOneByName() {
	name :=  u.Ctx.Input.Param(":name")
	if name != "" {
		filters := map[string]string { "name" : name }
		ob, err := (&modules.User{}).GetOneByFilter(filters)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = ob
		}
	}
	u.ServeJSON()
}

// @Title GetAll
// @Description find users
// @Success 200 []{user} []models.User
// @Failure 403 
// @router / [get]
func (u *UserController) GetAll() {
	ob, err := (&modules.User{}).GetAll()
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = ob
		}
	u.ServeJSON()
}

// @Title GetAll
// @Description find users
// @Success 200 []{user} []models.User
// @Failure 403 
// @router / [get]
func (u *UserController) Create() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	objectid, err :=  (&modules.User{}).Create(user)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = map[string]string{"ObjectId": objectid}
	}
	u.ServeJSON()
}
