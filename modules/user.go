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

type User struct{}

func (u *User)GetOneById(ObjectId string) (user models.User, err error){
	result := models.User{}
	data, err := (&tools.MongoHelper{}).GetOneById(collectionname, ObjectId)
	if err != nil {
		return models.User{}, err
	}
	err = bson.Unmarshal(data, &result)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}

func (u *User)GetOneByFilter(filters map[string]string) (user models.User, err error){
	result := models.User{}
	data, err := (&tools.MongoHelper{}).GetOneByFilter(collectionname, filters)
	if err != nil {
		return models.User{}, err
	}
	err = bson.Unmarshal(data, &result)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}

func (u *User)GetAll() (users []models.User,err error){
	result := []models.User{}
	data, err := (&tools.MongoHelper{}).GetAll(collectionname)
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

func (u *User)Create(user models.User) (objectId string ,err error) {
	objectId, err = (&tools.MongoHelper{}).Create(collectionname, user)
	if err != nil {
		log.Panicln(err)
		return "", err
	}
	return objectId, nil
}
