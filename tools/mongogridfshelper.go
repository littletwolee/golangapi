package tools

import(
//	"github.com/astaxie/beego"
//	"gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
	"log"
	"errors"
	"time"
	"golangapi/models"
)

const gridfscname = "fs"

type MongoGridFSHelper struct{}

func (m *MongoGridFSHelper) GetFileById (collectionname string, objectId string) (result []byte, err error){
	session := Session()
	db := session.DB(dbname)
	collection := db.GridFS(collectionname)
	data := bson.M{}
	if ! bson.IsObjectIdHex(objectId) {
		err = errors.New("It is not a objectId")
		log.Println(err)
		return nil, err
	}
	err = collection.Find(bson.M{"_id": bson.ObjectIdHex(objectId)}).One(&data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result, err = bson.Marshal(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}


func (m *MongoGridFSHelper) UploadFile (filemode models.Filemodel) (objectId string, err error) {
	session := Session()
	db := session.DB(dbname)
	newid := bson.NewObjectId()
	filename := newid.Hex() + "." + filemode.Filetype
	file, err := db.GridFS(gridfscname).Create(filename)
	if err != nil {
		log.Println(err)
		return "", err
	}
	file.SetId(newid)
	file.SetName(filename)
	file.SetUploadDate(time.Now())
	file.SetContentType(filemode.Contenttype)
	_, err = file.Write(filemode.Filedata)
	if err != nil {
		log.Println(err)
		return "", err
	}
	err = file.Close()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return newid.Hex(), nil
}


func (m *MongoGridFSHelper) DeleteFileById (objectId string) error {	
	session := Session()
	db := session.DB(dbname)
	err := db.GridFS(gridfscname).RemoveId(bson.ObjectIdHex(objectId))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
