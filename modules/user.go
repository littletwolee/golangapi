package modules

import(
	"mongoapi/models"
	"mongoapi/tools"
	"log"
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

func GetAll() (users []models.User,err error){
	result := []models.User{}
	data, err := tools.GetAll(collectionname)
	if err != nil {
		log.Panicln(err)
		return result, err
	}
	for _, item := range data {
		resultitem := models.User{}
		err = bson.Unmarshal(item, &resultitem)
		if err != nil {
			log.Panicln(err)
			return result, err
		}
		result = append(result, resultitem)
	} 
	return result, nil
}
