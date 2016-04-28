package modules

import(
//	"mongoapi/models"
	"mongoapi/tools"
//	"log"
	"gopkg.in/mgo.v2/bson"
)

var(
	versionrulecname       string = "versionrule"
)

type VersionRule struct{}

func (v *VersionRule) GetRuleByFilter (filters map[string]interface{}) (result bson.M, err error){
	data, err := (&tools.MongoHelper{}).GetFieldByFilter(versionrulecname, filters)
	err = bson.Unmarshal(data, &result)
	return result, err
}
