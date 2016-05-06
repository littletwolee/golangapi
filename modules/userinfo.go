package modules

import(
	"mongoapi/models"
	"mongoapi/tools"
	"log"
	"gopkg.in/mgo.v2/bson"
)
const userinfocname = "userinfo"
	

type Userinfo struct{}

func (u *Userinfo) GetOneUserinfoById (ObjectId string) (userinfo models.Userinfo, err error){
	result := models.Userinfo{}
	data, err := (&tools.MongoHelper{}).GetOneById(userinfocname, ObjectId)
	err = bson.Unmarshal(data, &result)
	return result, err
}

// func (u *Userinfo) GetOneUserinfoByFilter (filters map[string]string) (userinfo models.Userinfo, err error){
// 	result := models.Userinfo{}
// 	data, err := (&tools.MongoHelper{}).GetOneByFilter(userinfocname, filters)
// 	err = bson.Unmarshal(data, &result)
// 	return result, err
// }

func (u *Userinfo) CreateUserinfo (userinfo models.Userinfo) (objectId string ,err error) {
	return (&tools.MongoHelper{}).Create(userinfocname, userinfo)
}

func (u *Userinfo) DeleteUserinfo (objectId string) error {
	return (&tools.MongoHelper{}).DeleteDoc(userinfocname, objectId)
}

func (u *Userinfo) UpdateUserinfoById(ObjectId string, userinfo map[string]interface{}) error {
	return (&tools.MongoHelper{}).UpdateById(userinfocname, ObjectId, userinfo)
}

func (u *Userinfo) UploadUserPic(file []byte, parameters map[string]interface{}, thisnum int, maxnum) error{
	switch {
	case thisnum == 1 && thisnum == maxnum:
		objectId, err := (&tools.MongoGridFSHelper{}).UploadFile(file, parameters)
	case thisnum == 1 && thisnum < maxnum:
		(&tools.Filehelper{}).
	}
	if err != nil {
			log.Println(err)
			return err
		}
		userid := parameters["imname"].(string)
		userinfo := bson.M{"$set" : bson.M{"pic" : objectId}}
		err = (&tools.MongoHelper{}).UpdateById(userinfocname, userid, userinfo)
		if err != nil {
			log.Println(err)
			err = (&tools.MongoGridFSHelper{}).DeleteFileById(objectId)
			if err != nil {
				log.Println(err)
				return err
			}
			return err
		}
	return nil
}
