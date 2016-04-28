package controllers

import (
	"mongoapi/modules"
//	"mongoapi/models"
	"github.com/astaxie/beego"
//	"encoding/json"
//	"strconv" 
//	"log"
)

// Operations about object
type VersionRuleController struct {
	beego.Controller
}

// @Title GetOneVersionRuleByVersion
// @Description find rule by filters
// @Param	version	     "the version you want to get"
// @Success 200 string 
// @Failure 403 :version is empty
// @router /:version [get]
func (v *VersionRuleController) GetOneVersionRuleByVersion() {
	version :=  v.Ctx.Input.Param(":version")
	if version != "" {
		filters := make(map[string]interface{})
		filters["query"] = map[string]string{ "version" : version }
		filters["field"] = map[string]int{ "rule" : 1 , "_id" : 0 }
		ob, err := (&modules.VersionRule{}).GetRuleByFilter(filters)
		if err != nil {
			v.Data["json"] = err.Error()
		} else {
			v.Data["json"] = ob
		}
	}
	v.ServeJSON()
}
