package modules

import(
	"mongoapi/models"
	"mongoapi/tools"
//	"log"
	"gopkg.in/mgo.v2/bson"
)

var(
	collectionname       string = "user"
	
)

func GetOne(ObjectId string) models.User{
	result := models.User{}
	bson.Unmarshal(tools.GetOne(collectionname, ObjectId), &result)
	return result
}
