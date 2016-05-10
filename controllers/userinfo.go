package controllers

import (
	"mongoapi/modules"
	"mongoapi/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"strconv" 
	"log"
)

// Operations about object
type UserinfoController struct {
	beego.Controller
}

// @Title GetOneUserinfoById
// @Description find userinfo by objectid
// @Param	objectId	"the objectid you want to get"
// @Success 200 {userinfo} models.Userinfo
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (u *UserinfoController) GetOneUserinfoById() {
	objectId := u.Ctx.Input.Param(":objectId")
	if objectId != "" {
		ob, err := (&modules.Userinfo{}).GetOneUserinfoById(objectId)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = ob
		}
	}
	u.ServeJSON()
}

// @Title CreateUserinfo
// @Description create userinfo
// @Success 200 objectid
// @Failure 403 
// @router / [post]
func (u *UserinfoController) CreateUserinfo() {
	var userinfo models.Userinfo
	json.Unmarshal(u.Ctx.Input.RequestBody, &userinfo)
	objectid, err :=  (&modules.Userinfo{}).CreateUserinfo(userinfo)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = map[string]string{"ObjectId": objectid}
	}
	u.ServeJSON()
}

// @Title DeleteUserinfo
// @Description delete userinfo
// @Success 200 err nil
// @Failure 403 
// @router / [delete]
func (u *UserinfoController) DeleteUserinfo() {
	objectId := u.Ctx.Input.Param(":objectId")
	err := (&modules.Userinfo{}).DeleteUserinfo(objectId)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = map[string]string{"status": strconv.FormatBool(true)}
	}
	u.ServeJSON()
}

// @Title UpdateUserinfoById
// @Description update userinfo
// @Success 200 err nil
// @Failure 403 
// @router / [post]
func (u *UserinfoController) UpdateUserinfoById() {
	objectId := u.Ctx.Input.Param(":objectId")
	var userinfo map[string]interface{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &userinfo)
	err := (&modules.Userinfo{}).UpdateUserinfoById(objectId, userinfo)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = map[string]string{"status": strconv.FormatBool(true)}
	}
	u.ServeJSON()
}

// @Title UpdateUserinfoById
// @Description upload pic
// @Success 200 err nil
// @Failure 403 
// @router / [post]
func (u *UserinfoController) UploadUserPic() {
	var resultdata = new(models.ResponseResult)
	f, h, err := u.GetFile("file")
	if err != nil {
		resultdata = &models.ResponseResult{nil, err.Error(), 500, false}
	}
	f.Close()
	file := models.Filemodel{}
	file.Filename = h.Filename
	file.Contenttype = u.Input().Get("Content-Type")
	file.Filetype = u.Input().Get("filetype")
	bytesize ,err := strconv.Atoi(u.Input().Get("bytesize"))
	if err != nil {
		resultdata = &models.ResponseResult{nil, err.Error(), 500, false}
	}
	filedata := make([]byte, bytesize)
	_, err = f.Read(filedata)
	if err != nil {
		resultdata = &models.ResponseResult{nil, err.Error(), 500, false}
	}
	file.Filedata = filedata
	file.Currentchunk ,err = strconv.Atoi(u.Input().Get("currentchunk"))
	if err != nil {
		resultdata = &models.ResponseResult{nil, err.Error(), 500, false}
	}
	file.Maxchunks ,err = strconv.Atoi(u.Input().Get("maxchunks"))
	if err != nil {
		resultdata = &models.ResponseResult{nil, err.Error(), 500, false}
	}
	filename ,err := (&modules.Userinfo{}).UploadUserPic(file)
	if err != nil {
		resultdata = &models.ResponseResult{nil, err.Error(), 500, false}
	}
	resultdata = &models.ResponseResult{filename, err.Error(), 200, true}
	log.Println(resultdata.Filename)
	u.Data["json"] = resultdata
	u.ServeJSON()
}

