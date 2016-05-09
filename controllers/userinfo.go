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

	
	// file, fileheader, err := u.GetFile("pic")
	// if err != nil {
	// 	u.Data["json"] = err.Error()
	// } else {
	// 	path :="./" + fileheader.Filename
	// 	file.Close()
	// 	u.SaveToFile("pic", path)
	// 	u.Data["json"] = map[string]string{"status": strconv.FormatBool(true)}
	// }

	
// 	f, h, _ := this.GetFile("image")	//获取上传的文件
// 	path := this.Input().Get("url")	//存文件的路径    
// path = path[7:]	    
// path = "./static/img/" + path + "/" + h.Filename	    
// f.Close()	// 关闭上传的文件，不然的话会出现临时文件不能清除的情况    
// this.SaveToFile("image", path)	//存文件    WaterMark(path)	//给文件加水印    
// this.Redirect("/addphoto", 302)
	// objectId := u.Ctx.Input.Param(":objectId")
	// var userinfo map[string]interface{}
	// json.Unmarshal(u.Ctx.Input.RequestBody, &userinfo)
	// err := (&modules.Userinfo{}).UploadUserPic(objectId, userinfo)
	// if err != nil {
	// 	u.Data["json"] = err.Error()
	// } else {
	// 	u.Data["json"] = map[string]string{"status": strconv.FormatBool(true)}
	// }
	f, _, _ := u.GetFile("file")
	//rule := u.Input().Get("rule")
	// //获取上传的文件
	// path := u.Input().Get("url")	//存文件的路径
	// log.Println(path)
	// path = path[7:]	    
	//path := "./"+ h.Filename
	 f.Close()	// 关闭上传的文件，不然的话会出现临时文件不能清除的情况    
	//u.SaveToFile("file", path)	//存文件    WaterMark(path)	//给文件加水印
	data := make([]byte, 8668)
	_, err := f.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data[0])
	u.ServeJSON()
}

