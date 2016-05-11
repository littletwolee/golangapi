package controllers

import (
	"golangapi/modules"
	"golangapi/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"strconv" 
//	"log"
)

// Operations about object
type UserController struct {
	beego.Controller
}

// @Title GetOneUserById
// @Description find user by objectid
// @Param	objectId	"the objectid you want to get"
// @Success 200 {user} models.User
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (u *UserController) GetOneUserById() {
	objectId := u.Ctx.Input.Param(":objectId")
	if objectId != "" {
		ob, err := (&modules.User{}).GetOneUserById(objectId)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = ob
		}
	}
	u.ServeJSON()
}

// @Title GetOneUserByName
// @Description find user by filters
// @Param	name	"the name you want to get"
// @Success 200 {user} models.User
// @Failure 403 :name is empty
// @router /:name [get]
func (u *UserController) GetOneUserByName() {
	name :=  u.Ctx.Input.Param(":name")
	if name != "" {
		filters := map[string]string { "name" : name }
		ob, err := (&modules.User{}).GetOneUserByFilter(filters)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = ob
		}
	}
	u.ServeJSON()
}

// @Title GetAllUsers
// @Description find users
// @Success 200 []{user} []models.User
// @Failure 403 
// @router / [get]
func (u *UserController) GetAllUsers() {
	ob, err := (&modules.User{}).GetAllUsers()
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = ob
		}
	u.ServeJSON()
}

// @Title CreateUser
// @Description create user
// @Success 200 objectid
// @Failure 403 
// @router / [post]
func (u *UserController) CreateUser() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	objectid, err :=  (&modules.User{}).CreateUser(user)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = map[string]string{"ObjectId": objectid}
	}
	u.ServeJSON()
}

// @Title DeleteUser
// @Description delete user
// @Success 200 err nil
// @Failure 403 
// @router / [delete]
func (u *UserController) DeleteUser() {
	objectId := u.Ctx.Input.Param(":objectId")
	err := (&modules.User{}).DeleteUser(objectId)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = map[string]string{"status": strconv.FormatBool(true)}
	}
	u.ServeJSON()
}
