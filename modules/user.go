package modules

import(
	"mongoapi/models"
	"mongoapi/tools"
)

var(
	collectionname       string = "user"
	
)

func GetOne(ObjectId string) models.User{
	return tools.GetOne(collectionname, ObjectId).(models.User)
}
