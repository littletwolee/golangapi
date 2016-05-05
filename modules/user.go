package modules

import(
	"mongoapi/models"
	"mongoapi/tools"
	"log"
	"gopkg.in/mgo.v2/bson"
)

const usercname = "user"

type User struct{}

func (u *User) GetOneUserById (ObjectId string) (user models.User, err error){
	result := models.User{}
	data, err := (&tools.MongoHelper{}).GetOneById(usercname, ObjectId)
	err = bson.Unmarshal(data, &result)
	return result, err
}

func (u *User) GetOneUserByFilter (filters map[string]string) (user models.User, err error){
	result := models.User{}
	data, err := (&tools.MongoHelper{}).GetOneByFilter(usercname, filters)
	err = bson.Unmarshal(data, &result)
	return result, err
}

func (u *User) GetAllUsers () (users []models.User,err error){
	result := []models.User{}
	data, err := (&tools.MongoHelper{}).GetAll(usercname)
	for _, item := range data {
		resultitem := models.User{}
		err = bson.Unmarshal(item, &resultitem)
		if err != nil {
			log.Panicln(err)
			return result, err
		}
		result = append(result, resultitem)
	} 
	return result, err
}

func (u *User) CreateUser (user models.User) (objectId string ,err error) {
	return (&tools.MongoHelper{}).Create(usercname, user)
}

func (u *User) DeleteUser (objectId string) error {
	return (&tools.MongoHelper{}).DeleteDoc(usercname, objectId)
}
