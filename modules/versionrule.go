package modules

import(
	"mongoapi/models"
	"mongoapi/tools"
//	"log"
	"gopkg.in/mgo.v2/bson"
)

var(
	versionrulecname       string = "versionrule"
)

type VersionRule struct{}

func (v *VersionRule) GetRuleByFilter (filters []map[string]interface{}) (resule []int, err error){
	data, err := (&tools.MongoHelper{}).GetFieldByFilter(versionrulecname, filters)
	versionrule := models.VersionRule{}
	err = bson.Unmarshal(data, &versionrule)
	resule = versionrule.Rule
	return resule, err
}
