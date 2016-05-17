package modules

import(
	"golangapi/models"
	"golangapi/tools"
//	"strconv"
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

func (u *Userinfo) UploadUserPic(filemode models.Filemodel, userid string, userpic string) (string, error) {
	if fileid, err := (&tools.Filehelper{}).UploadFileToMongo(filemode); err == nil{
		if (filemode.Currentchunk == 0 &&
			filemode.Currentchunk == filemode.Maxchunks - 1) ||
			filemode.Currentchunk == filemode.Maxchunks - 1 {
			userinfo := bson.M{"pic" : fileid}
			err = (&tools.MongoHelper{}).UpdateById(userinfocname, userid, userinfo)
			if err == nil {
				fileid = userpic
			}
			if userpic != "" {
				err = (&tools.MongoGridFSHelper{}).DeleteFileById(fileid)
				if err != nil {
					return "", err
				}
				return fileid, nil
			}
			return "", err
		}
		return fileid, nil
	} else {
		return "", err
	}
	
}

func (u *Userinfo) DownloadUserPic(userpic string) (interface{}, error) {
	rangemode, err := (&tools.MongoGridFSHelper{}).GetFileById(userinfocname, userpic)
	if err != nil {
		return nil, err
	}
	// filename := userpic + "_" +
	// 	strconv.FormatInt(rangemode.(models.Rangemodel).Start, 10) + "_" +
	// 	strconv.FormatInt(rangemode.(models.Rangemodel).End, 10)
	// if filechunkdata ,err := redishelper.GetVByK(filename, "bytes"); err == nil {
	// 			filemode.Filedata = append(filechunkdata.([]byte), filemode.Filedata...)
	// 		}
	// if err := (&tools.RedisHelper{}).SetKVBySETEX(filename, rangemode.(models.Rangemodel).Filedata, 60); err == nil {
	// 	return nil, err
	// }
	return rangemode, nil
}
