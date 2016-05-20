package controllers

import (
	"golangapi/modules"
	"golangapi/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"strconv"
//	"log"
	"strings"
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
	u.ServeJSON(true)
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
// @router / [form]
func (u *UserinfoController) UploadUserPic() {
	resultdata := &models.ResponseResult{}
	f, h, err := u.GetFile("file")
	if err != nil {
		resultdata = &models.ResponseResult{nil, err.Error(), 500, false}
	}
	f.Close()
	file := models.Filemodel{}
	file.Filename = strings.Trim(h.Filename," ")
	file.Contenttype = h.Header.Get("Content-Type")
	file.Filetype = u.Input().Get("filetype")
	userid := u.Input().Get("userid")
	userpic := strings.Trim(u.Input().Get("userpic")," ")
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
	filename ,err := (&modules.Userinfo{}).UploadUserPic(file, userid, userpic)
	if err != nil {
		resultdata = &models.ResponseResult{nil, err.Error(), 500, false}
	}
	resultdata = &models.ResponseResult{filename, "", 200, true}
	u.Data["json"] = resultdata
	u.ServeJSON()
}

// @Title DownloadUserPic
// @Description download user pic
// @Success 200 err nil
// @Failure 403 
// @router / [get]
func (u *UserinfoController) DownloadUserPic() {
	userpic := u.Ctx.Input.Param(":userpic")
	result, err := (&modules.Userinfo{}).DownloadUserPic(userpic)
	if err != nil {
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	}
	rangemode := result.(*models.Rangemodel)
	u.Ctx.Output.Header("Content-Type", rangemode.Contenttype)
	u.Ctx.Output.Header("Content-Length", strconv.FormatInt(rangemode.Size, 10))
	u.Ctx.Output.Body(rangemode.Filedata)
}

// @Title DownBigFile
// @Description download big file
// @Success 200 err nil
// @Failure 403 
// @router / [get]
// func (u *UserinfoController) DownBigFile() {
// 	fileid := u.Ctx.Input.Param(":fileid") 
// 	result, err := (&modules.Userinfo{}).DownloadBigFile(fileid)
// 	if err != nil {
// 		u.Data["json"] = err.Error()
// 		u.ServeJSON()
// 		return
// 	}
// 	rangemode := result.(*models.Rangemodel)
// 	if rangestr := u.Ctx.Input.Header("range"); rangestr == "" {
// 		u.Ctx.Output.Header("Content-Type", rangemode.Contenttype)
// 		u.Ctx.Output.Header("Content-Length", strconv.FormatInt(rangemode.End, 10))
// 		u.Ctx.Output.Header("Accept-Ranges", "bytes")
// 		u.Ctx.Output.Body(rangemode.Filedata)
// 		u.Ctx.Output.SetStatus(200)
// 	} else {
// 		if start, end, err := tools.SplitRange(rangestr); err != nil {
// 			u.Data["json"] = err.Error()
// 			u.ServeJSON()
// 			return 
// 		} else {
// 			if start >= rangemode.Size || end >= rangemode.Size {
// 				u.Ctx.Output.Header("Content-Range", "bytes */" + strconv.FormatInt(rangemode.Size, 10))
// 				u.Ctx.Output.SetStatus(416)
// 				return
// 			}
// 			u.Ctx.Output.Header("Content-Range", "bytes " +
// 				strconv.FormatInt(start, 10) + "-" +
// 				strconv.FormatInt(end, 10)+ "/" +
// 				strconv.FormatInt(rangemode.Size, 10))
// 			if start == end {
// 				u.Ctx.Output.Header("Content-Length", "0")
// 			} else {
// 				u.Ctx.Output.Header("Content-Length", strconv.FormatInt(end - start + 1, 10))}
// 			u.Ctx.Output.Header("Content-Type", rangemode.Contenttype)
// 			u.Ctx.Output.Body(rangemode.Filedata[start : end + 1])
// 			u.Ctx.Output.Header("Accept-Ranges", "bytes")
// 			u.Ctx.Output.SetStatus(206)
// 			u.Ctx.Output.Header("Cache-Control", "no-cache")}
// 	}
// }

// @Title CreateRelationship
// @Description add friend
// @Success 200 err nil
// @Failure 403 
// @router / [post]
func(u *UserinfoController) CreateRelationship() {
	relationship := models.Relationship{ -1, -1}
	json.Unmarshal(u.Ctx.Input.RequestBody, &relationship)
	if relationship.Userid == -1 || relationship.Friendid == -1 {
		u.Data["json"] = map[string]string{"status1": strconv.FormatBool(false)}
		u.ServeJSON()
		return
	}
	// userid, err := strconv.Atoi(relationship.Userid)
	// if err != nil {
	// 	u.Data["json"] = err.Error()
	// 	u.ServeJSON()
	// 	return
	// }
	// friendid, err := strconv.Atoi(relationship.Friendid)
	// if err != nil {
	// 	u.Data["json"] = err.Error()
	// 	u.ServeJSON()
	// 	return
	// }
	err := (&modules.Userinfo{}).CreateRelationship(relationship.Userid, relationship.Friendid)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = map[string]string{"status": strconv.FormatBool(true)}
	}
	u.ServeJSON()
}
