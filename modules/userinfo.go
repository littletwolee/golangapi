package modules

import(
	"golangapi/models"
	"golangapi/tools"
//	"log"
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

func (u *Userinfo) UploadUserPic(filemode models.Filemodel) (string, error) {
	// if 
	// switch {
	// case parameters["offset"] == 1 && thisnum == maxnum:
	// 	objectId, err := (&tools.MongoGridFSHelper{}).UploadFile(file, parameters)
	// case thisnum == 1 && thisnum < maxnum:
	// 	(&tools.Filehelper{}).
	// }
	// if err != nil {
	// 		log.Println(err)
	// 		return err
	// 	}
	// 	userid := parameters["imname"].(string)
	// 	userinfo := bson.M{"$set" : bson.M{"pic" : objectId}}
	// 	err = (&tools.MongoHelper{}).UpdateById(userinfocname, userid, userinfo)
	// 	if err != nil {
	// 		log.Println(err)
	// 		err = (&tools.MongoGridFSHelper{}).DeleteFileById(objectId)
	// 		if err != nil {
	// 			log.Println(err)
	// 			return err
	// 		}
	// 		return err
	// 	}
	// return nil


	
	filename := ""
	filehelper := &tools.Filehelper{}
	redishelper := &tools.RedisHelper{}
	if filemode.Currentchunk == 0 {
		filename = tools.GetGuid() 
	} else {
		filename = filemode.Filename
	}
	switch {
	case filemode.Currentchunk == 0 && filemode.Currentchunk < filemode.Maxchunks - 1:
		if err := redishelper.SetKVBySETEX(filename, filemode.Filedata, 300); err == nil {
			return filename, nil
		}
		
		//return filehelper.WriteFile(filename, filemode.Filedata)
		
	case filemode.Currentchunk == 0 && filemode.Currentchunk == filemode.Maxchunks - 1 :
		return filehelper.WriteFile(filename, filemode.Filedata)
	case filemode.Currentchunk != 0 && filemode.Currentchunk < filemode.Maxchunks - 1 :
		if filechunkdata ,err := redishelper.GetVByK(filename, "bytes"); err == nil {
				return filename, nil
		} else {
			newdata := append(filechunkdata.([]byte), filemode.Filedata...)
			return filehelper.WriteFile(filename, newdata)
		}
	}
	return "", (&tools.ResultHelp{}).NewErr("server err")
}
