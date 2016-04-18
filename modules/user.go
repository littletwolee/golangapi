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

func GetOne(ObjectId string) (user models.User, err error){
	result := models.User{}
	data, err := tools.GetOne(collectionname, ObjectId)
	if err != nil {
		return models.User{}, err
	}
	err = bson.Unmarshal(data, &result)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}

func GetAll() map[string]models.User{
	result := map[string]models.User{}
	bson.Unmarshal(tools.GetAll(collectionname), &result)
	return result
}
