package modules

import(
//	"mongoapi/models"
	"mongoapi/tools"
)

var(
	collectionname       string = "user"
	
)

func GetOne(ObjectId string) interface{}{
	return tools.GetOne(collectionname, ObjectId)	
}
