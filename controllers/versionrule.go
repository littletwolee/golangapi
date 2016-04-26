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
// @Param	name	"the name you want to get"
// @Success 200 []int 
// @Failure 403 :version is empty
// @router /:version [get]
func (v *VersionRuleController) GetOneVersionRuleByVersion() {
	version :=  v.Ctx.Input.Param(":version")
	if version != "" {
		var filters []map[string]interface{}
		filters = append(filters, map[string]interface{}{ "version" : version })
		filters = append(filters, map[string]interface{}{ "rule" : 1 })
		ob, err := (&modules.VersionRule{}).GetRuleByFilter(filters)
		if err != nil {
			v.Data["json"] = err.Error()
		} else {
			v.Data["json"] = ob
		}
	}
	v.ServeJSON()
}
